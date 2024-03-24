package micro

import "google.golang.org/grpc"

type Abstraction interface {
	Register(srv interface{}, desc grpc.ServiceDesc)
	Deregister()
	CreateLease()
	WithRetryBefore(func())
	WithRetryAfter(func())

	Watcher(config *[]string, service *map[string][]string, http *map[string]string)
}
