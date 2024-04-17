package grpc

import (
	"google.golang.org/grpc/resolver"
)

// 自定义name resolver

const (
	myScheme   = "grpc"
	myEndpoint = "resolver.lhdht.com"
)

var addr = []string{"127.0.0.1:8972", "127.0.0.1:8973"}

// grpcResolver 自定义name resolver，实现Resolver接口
type grpcResolver struct {
	target    resolver.Target
	cc        resolver.ClientConn
	addrStore map[string][]string
}

func (r *grpcResolver) ResolveNow(o resolver.ResolveNowOptions) {
	addrStrs := r.addrStore[r.target.Endpoint]
	addrList := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrList[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrList})
}

func (*grpcResolver) Close() {}

// grpcResolverBuilder 需实现 Builder 接口
type grpcResolverBuilder struct{}

func (*grpcResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &grpcResolver{
		target: target,
		cc:     cc,
		addrStore: map[string][]string{
			myEndpoint: addr,
		},
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}
func (*grpcResolverBuilder) Scheme() string { return myScheme }

func init() {
	// 注册 q1miResolverBuilder
	resolver.Register(&grpcResolverBuilder{})
}
