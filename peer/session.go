package peer

import (
	"net"
	"sync"
	"tyszj"
)

type tcpSession struct {
	PeerBundle
	//原始连接
	conn net.Conn
	//退出同步器,收发全部完成
	exitSync sync.WaitGroup
	//发送队列
	sendQueue tyszj.IPipe
	//清除数据锁
	cleanUpGuard sync.Mutex

	id int64

	peer tyszj.IPeer
}

func (self *tcpSession) Peer() tyszj.IPeer {
	return self.peer
}

func (self *tcpSession) ID() int64 {
	return self.id
}

func (self *tcpSession) SetID(id int64) {
	self.id = id
}

func (self *tcpSession) Raw() interface{} {
	return self.conn
}

func (self *tcpSession) Close() {
	self.sendQueue.Add(nil)
}

func (self *tcpSession) Send(msg interface{}) {
	if nil == msg {
		return
	}
	self.sendQueue.Add(msg)
}

func (self *tcpSession) ReadMessage() (msg interface{}, err error) {
	return
}

func (self *tcpSession) recvLoop() {
	for self.conn != nil {
		msg, err := self.ReadMessage()
		if err != nil {
			self.Close()
			self.PostEvent(&tyszj.RecvMsgEvent{self, &tyszj.SessionClosed{}})
			break
		}
		self.PostEvent(&tyszj.RecvMsgEvent{self, msg})
	}
}
func (self *tcpSession) sendLoop() {

}

func (self *tcpSession) cleanUp() {

}

func (self *tcpSession) Start() {

}

func newSession(conn net.Conn, peer tyszj.IPeer) *tcpSession {
	ses := &tcpSession{
		conn:      conn,
		sendQueue: tyszj.NewPipe(),
		peer:      peer,
	}
	return ses
}
