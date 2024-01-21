package micro

type Abstraction interface {
	Register(service string)
	Deregister()
	CreateLease()
	WithRetryBefore(func())
	WithRetryAfter(func())

	Watcher(config *[]string, service *map[string][]string)
}
