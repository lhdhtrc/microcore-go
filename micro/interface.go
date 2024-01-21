package micro

type Abstraction interface {
	Register(service string)
	Deregister()
	CreateLease(retry func())

	Watcher(config *[]string, service *map[string][]string)
}
