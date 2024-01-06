package micro

type Abstraction interface {
	Register(service string)
	Deregister()
	CreateLease()

	Watcher(config *[]string, service *map[string][]string)
}
