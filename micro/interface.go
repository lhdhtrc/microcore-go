package micro

type Abstraction interface {
	Register(prefix string, srv interface{})
	Deregister()
	CreateLease()
	WithRetryBefore(func())
	WithRetryAfter(func())

	Watcher(config *[]string, service *map[string][]string)
}
