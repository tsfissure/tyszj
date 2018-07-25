package peer

import (
	"tyszj"
)

type PeerProperty struct {
	queue tyszj.IEventQueue
	addr  string
}

func (self *PeerProperty) Queue() tyszj.IEventQueue {
	return self.queue
}

func (self *PeerProperty) Address() string {
	return self.addr
}

func (self *PeerProperty) SetQueue(q tyszj.IEventQueue) {
	self.queue = q
}

func (self *PeerProperty) SetAddress(addr string) {
	self.addr = addr
}
