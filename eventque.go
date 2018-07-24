package tyszj

import (
	"sync"
)

type IEventQueue interface {
	//开始循环
	StartLoop()

	//手动结束循环
	StopLoop()

	//等待退出
	Wait()

	//post event
	Post(callback func())
}

type eventQueue struct {
	queue chan func()

	endSignal sync.WaitGroup
}

func (self *eventQueue) Post(callback func()) {
	if nil == callback {
		return
	}
	self.queue <- callback
}

func (self *eventQueue) StartLoop() {
	self.endSignal.Add(1)
	go func() {
		for callback := range self.queue {
			if nil == callback {
				break
			}
			callback()
		}
		self.endSignal.Done()
	}()
}

func (self *eventQueue) StopLoop() {
	self.queue <- nil
}

func (self *eventQueue) Wait() {
	self.endSignal.Wait()
}

func SessionQueuedCall(ses ISession, callback func()) {
	if nil == ses {
		return
	}
	q := ses.Peer().(IPeerProperty).Queue()
	QueuedCall(q, callback)
}
func QueuedCall(queue IEventQueue, callback func()) {
	if nil == queue {
		callback()
	} else {
		queue.Post(callback)
	}
}

const cDefaultQueueSize = 100

func NewEventQueue() IEventQueue {
	return &eventQueue{
		queue: make(chan func(), cDefaultQueueSize),
	}
}
