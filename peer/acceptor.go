package peer

import (
	"fmt"
	"net"
	"tyszj"
)

type tcpActepter struct {
	SessionManager
	PeerBundle

	queue tyszj.IEventQueue
	addr  string

	listener net.Listener
	RunningTag
}

func (self *tcpActepter) Queue() tyszj.IEventQueue {
	return self.queue
}
func (self *tcpActepter) Address() string {
	return self.addr
}

func (self *tcpActepter) Start() {
	self.WaitStopFinished()
	if self.IsRunning() {
		return
	}
	ln, err := net.Listen("tcp", self.Address())
	if err != nil {
		fmt.Printf("#tcp.listen FAIL[%s] #v", self.Address(), err.Error())
		self.SetRunningStat(false)
		return
	}
	self.listener = ln
	fmt.Printf("#tcp.listen(%s)", self.Address())
	go self.accept()
}
func (self *tcpActepter) Stop() {
	if !self.IsRunning() {
		return
	}
	if self.IsStopping() {
		return
	}
	self.StartStopping()
	self.listener.Close()
	self.CloseAllSession()
	self.WaitStopFinished()
}

func (self *tcpActepter) accept() {
	self.SetRunningStat(true)
	for {
		if self.IsStopping() {
			break
		}
		conn, err := self.listener.Accept()
		if err != nil {
			break
		}
		go self.onNewSession(conn)
	}
	self.SetRunningStat(false)
}
func (self *tcpActepter) onNewSession(conn net.Conn) {
	ses := newSession(conn, self)
	ses.Start()
	self.PostEvent(&tyszj.RecvMsgEvent{ses, &tyszj.SessionAccepted{}})
}
