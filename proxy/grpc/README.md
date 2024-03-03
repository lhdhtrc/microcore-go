## Lhdht Microservice Golang Version

## example
```go
package main

import (
	"context"
	grpcProxy "github.com/lhdhtrc/microservice-go/proxy/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func init() {
	encoding.RegisterCodec(grpcProxy.Codec())
}

func main() {
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		// 服务发现的服务列表
		endPoint := make(map[string][]string)

		// 根据fullMethodName获取可用节点
		nodes, ok := endPoint[fullMethodName]

		var cc *grpc.ClientConn
		err := status.Errorf(codes.Unimplemented, "Unknown method")

		if ok && len(nodes) != 0 {
			md, _ := metadata.FromIncomingContext(ctx)
			ctx = metadata.NewOutgoingContext(ctx, md.Copy())

			cc, err = grpc.DialContext(ctx, nodes[0])
		}

		return ctx, cc, err
	}

	grpc.WithDefaultCallOptions(grpc.CallContentSubtype(grpcProxy.Codec().Name()))
	server := grpc.NewServer(grpc.UnknownServiceHandler(grpcProxy.TransparentHandler(director)))
}
```