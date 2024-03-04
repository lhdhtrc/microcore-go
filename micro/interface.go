package micro

import "google.golang.org/grpc"

type Abstraction interface {
	Register(prefix string, srv interface{}, desc grpc.ServiceDesc)
	Deregister()
	CreateLease()
	WithRetryBefore(func())
	WithRetryAfter(func())

	Watcher(config *[]string, service *map[string][]string)
}
