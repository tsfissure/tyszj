package peer

import (
	"fmt"
	"net"
	"sync"
	"tyszj"
)

type tcpConnector struct {
	SessionManager
	PeerBundle
	RunningTag
	PeerProperty
	session   *tcpSession
	endSignal sync.WaitGroup
}

func (self *tcpConnector) Start() {
	self.WaitStopFinished()
	if self.IsRunning() {
		return
	}
	go self.connect(self.Address())
}

func (self *tcpConnector) Stop() {
	if !self.IsRunning() {
		return
	}
	if self.IsStopping() {
		return
	}
	self.StartStopping()
	self.session.Close()
	self.WaitStopFinished()
}

func (self *tcpConnector) connect(addr string) {
	self.SetRunningStat(true)
	conn, err := net.Dial("tcp", addr)
	self.session.conn = conn
	if err != nil {
		fmt.Println("#tcp.connect FAIL(%s) id(%d)", addr, self.session.id)
		return
	}
	self.endSignal.Add(1)
	self.session.Start()
	self.endSignal.Wait()
	self.session.conn = nil
	self.SetRunningStat(false)
}

func (self *tcpConnector) TypeName() string {
	return "tcp.Connector"
}

func init() {
	tyszj.RegisterPeerCreator(func() tyszj.IPeer {
		p := &tcpConnector{}
		return p
	})
}
