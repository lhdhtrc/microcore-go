package micro

type RegisterAbstraction interface {
	Register(service string)
	Deregister()
	CreateLease()
}

type DiscoverAbstraction interface {
	Watcher(config *[]DiscoverEntity)
}
