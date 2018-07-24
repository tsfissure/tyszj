package tyszj

import (
	"sync"
)

// pipe 用来发送和接收数据的队列
type IPipe interface {
	Add(msg interface{})
	Reset()
	Pick(retList *[]interface{}) (exit bool)

	Pipe() IPipe
}

type pipe struct {
	list []interface{}

	listGuard sync.Mutex
	listCond  *sync.Cond
}

func (self *pipe) Pipe() IPipe {
	return self
}

func (self *pipe) Add(msg interface{}) {
	self.listGuard.Lock()
	self.list = append(self.list, msg)
	self.listGuard.Unlock()

	self.listCond.Signal()
}

func (self *pipe) Reset() {
	self.list = self.list[0:0]
}

func (self *pipe) Pick(retList *[]interface{}) (exit bool) {
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
