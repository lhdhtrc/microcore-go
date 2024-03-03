package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

var (
	clientStreamDescForProxying = &grpc.StreamDesc{
		ServerStreams: true,
		ClientStreams: true,
	}
)

// RegisterService 注册代理服务
func RegisterService(server *grpc.Server, director StreamDirector, serviceName string, methodNames ...string) {
	streamer := &handler{director}
	fakeDesc := &grpc.ServiceDesc{
		ServiceName: serviceName,
		HandlerType: (*interface{})(nil),
	}
	for _, m := range methodNames {
		streamDesc := grpc.StreamDesc{
			StreamName:    m,
			Handler:       streamer.handler,
			ServerStreams: true,
			ClientStreams: true,
		}
		fakeDesc.Streams = append(fakeDesc.Streams, streamDesc)
	}
	server.RegisterService(fakeDesc, streamer)
}

// TransparentHandler 提供一个透明代理的方式，将未知的服务请求转发到后端服务，返回处理程序可以作为grpc.UnknownServiceHandler使用
func TransparentHandler(director StreamDirector) grpc.StreamHandler {
	streamer := &handler{director: director}
	return streamer.handler
}

type handler struct {
	director StreamDirector
}

/*
** handle 是代理功能的核心所在，负责将请求转发到其他服务。
** 就像调用任何 gRPC 服务端流一样。
** 使用 emptypb.Empty 类型作为服务器，在输入流和输出流之间转发调用。
 */
func (h *handler) handler(srv interface{}, serverStream grpc.ServerStream) error {
	fullMethodName, ok := grpc.MethodFromServerStream(serverStream)
	if !ok {
		// lowLevelServerStream 在上下文中不存在
		return status.Errorf(codes.Internal, "lowLevelServerStream not exists in context")
	}
	// director返回的上下文继承自serverStream.Context()
	outgoingCtx, backendConn, err := h.director(serverStream.Context(), fullMethodName)
	if err != nil {
		return err
	}

	clientCtx, clientCancel := context.WithCancel(outgoingCtx)
	defer clientCancel()
	// TODO(mwitkow): Add a `forwarded` header to metadata, https://en.wikipedia.org/wiki/X-Forwarded-For.
	clientStream, err := grpc.NewClientStream(clientCtx, clientStreamDescForProxying, backendConn, fullMethodName)
	if err != nil {
		return err
	}

	s2cErrChan := h.forwardServerToClient(serverStream, clientStream)
	c2sErrChan := h.forwardClientToServer(clientStream, serverStream)
	// 使用 select 语句进行非阻塞式等待, 避免程序陷入等待特定通道可读的死循环中。
	for i := 0; i < 2; i++ {
		select {
		case s2cErr := <-s2cErrChan:
			if s2cErr == io.EOF {
				if cCloseErr := clientStream.CloseSend(); cCloseErr != nil {
					return cCloseErr
				}
			} else {
				clientCancel()
				return status.Errorf(codes.Internal, "failed proxying s2c: %v", s2cErr)
			}
		case c2sErr := <-c2sErrChan:
			serverStream.SetTrailer(clientStream.Trailer())
			if c2sErr != io.EOF {
				return c2sErr
			}
			return nil
		}
	}
	return status.Errorf(codes.Internal, "gRPC proxying should never reach this stage.")
}

func (h *handler) forwardClientToServer(src grpc.ClientStream, dst grpc.ServerStream) chan error {
	ret := make(chan error, 1)
	go func() {
		f := &emptypb.Empty{}
		for i := 0; ; i++ {
			if err := src.RecvMsg(f); err != nil {
				ret <- err
				break
			}
			if i == 0 {
				// 这是一个有点取巧的做法，但客户端到服务器的头部信息只能在收到第一个客户端消息后才能读取，
				// 但必须在刷新第一个消息之前写入服务器流。
				// 这是唯一一个合适的地方来处理它。
				md, err := src.Header()
				if err != nil {
					ret <- err
					break
				}
				if err := dst.SendHeader(md); err != nil {
					ret <- err
					break
				}
			}
			if err := dst.SendMsg(f); err != nil {
				ret <- err
				break
			}
		}
	}()
	return ret
}

func (h *handler) forwardServerToClient(src grpc.ServerStream, dst grpc.ClientStream) chan error {
	ret := make(chan error, 1)
	go func() {
		f := &emptypb.Empty{}
		for i := 0; ; i++ {
			if err := src.RecvMsg(f); err != nil {
				ret <- err
				break
			}
			if err := dst.SendMsg(f); err != nil {
				ret <- err
				break
			}
		}
	}()
	return ret
}
