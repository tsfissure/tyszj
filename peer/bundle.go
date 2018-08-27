package peer

import (
	"tyszj"
)

type PeerBundle struct {
	callback tyszj.FEventCallback
}

func (self *PeerBundle) GetBundle() *PeerBundle {
	return self
}

func (self *PeerBundle) SetCallback(cb tyszj.FEventCallback) {
	self.callback = cb
}

func (self *PeerBundle) PostEvent(ev tyszj.IEvent) {
	if self.callback != nil {
		tyszj.SessionQueuedCall(ev.Session(), func() {
			self.callback(ev)
		})
	}
}
