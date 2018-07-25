package tyszj

type IPeer interface {
	Start()
	Stop()
	TypeName() string
}

type IPeerProperty interface {
	Queue() IEventQueue
	Address() string
	SetAddress(v string)
	SetQueue(v IEventQueue)
}

type GenericPeer interface {
	IPeer
	IPeerProperty
}
