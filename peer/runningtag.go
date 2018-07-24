package peer

import (
	"sync/atomic"
)

type RunningTag struct {
	running int32

	stopping chan bool
}

func (self *RunningTag) IsRunning() bool {
	return atomic.LoadInt32(&self.running) != 0
}

func (self *RunningTag) SetRunningStat(v bool) {
	if v {
		atomic.StoreInt32(&self.running, 1)
	} else {
		atomic.StoreInt32(&self.running, 0)
	}
}

func (self *RunningTag) WaitStopFinished() {
	if self.stopping != nil {
		<-self.stopping
		self.stopping = nil
	}
}

func (self *RunningTag) IsStopping() bool {
	return self.stopping != nil
}

func (self *RunningTag) StartStopping() {
	self.stopping = make(chan bool)
}
