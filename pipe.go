package tyszj

import (
	"sync"
)

// pipe 用来发送和接收数据的队列
type IPipe interface {
	Add(*BasicMessage)
	Reset()
	Pick(retList *[]*BasicMessage) (exit bool)

	Pipe() IPipe
}

type pipe struct {
	list []*BasicMessage

	listGuard sync.Mutex
	listCond  *sync.Cond
}

func (self *pipe) Pipe() IPipe {
	return self
}

func (self *pipe) Add(msg *BasicMessage) {
	self.listGuard.Lock()
	self.list = append(self.list, msg)
	self.listGuard.Unlock()

	self.listCond.Signal()
}

func (self *pipe) Reset() {
	self.list = self.list[0:0]
}

func (self *pipe) Pick(retList *[]*BasicMessage) (exit bool) {
	self.listGuard.Lock()

	if len(self.list) == 0 {
		self.listCond.Wait()
	}
	self.listGuard.Unlock()

	self.listGuard.Lock()
	for _, data := range self.list {
		if nil == data {
			exit = true
			break
		}
		*retList = append(*retList, data)
	}
	self.Reset()
	self.listGuard.Unlock()
	return
}

func NewPipe() IPipe {
	p := &pipe{}
	p.listCond = sync.NewCond(&p.listGuard)

	return p.Pipe()
}
