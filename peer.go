package tyszj

type IPeer interface {
	Start()
	Stop()
}

type IPeerProperty interface {
	Queue() IEventQueue
}
