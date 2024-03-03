package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewProxy(dst *grpc.ClientConn, opts ...grpc.ServerOption) *grpc.Server {
	opts = append(opts, DefaultProxyOpt(dst))
	return grpc.NewServer(opts...)
}

func DefaultProxyOpt(cc *grpc.ClientConn) grpc.ServerOption {
	return grpc.UnknownServiceHandler(TransparentHandler(DefaultDirector(cc)))
}

func DefaultDirector(cc *grpc.ClientConn) StreamDirector {
	return func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		ctx = metadata.NewOutgoingContext(ctx, md.Copy())
		return ctx, cc, nil
	}
}
