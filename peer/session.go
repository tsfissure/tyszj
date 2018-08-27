package peer

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync"
	"tyszj"
)

type tcpSession struct {
	*PeerBundle
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

func (self *tcpSession) Send(msg tyszj.BasicMessage) {
	self.sendQueue.Add(&msg)
}

var (
	ErrMinPacket    = errors.New("packet short size")
	ErrValidMsgSize = errors.New("valid packet size")
)

func (self *tcpSession) readMessage() (msg *tyszj.BasicMessage, err error) {
	reader, ok := self.Raw().(io.Reader)
	if !ok || reader == nil {
		return nil, nil
	}
	var sizeBuf = make([]byte, 2)
	_, err = io.ReadFull(reader, sizeBuf)
	if err != nil {
		return nil, ErrMinPacket
	}
	if len(sizeBuf) < 2 {
		return nil, ErrMinPacket
	}
	size := binary.LittleEndian.Uint16(sizeBuf)
	if size > 60000 || size < 2 {
		return nil, ErrValidMsgSize
	}
	body := make([]byte, size)
	_, err = io.ReadFull(reader, body)
	if err != nil {
		return
	}
	msgid := binary.LittleEndian.Uint16(body[:2])
	msg = &tyszj.BasicMessage{msgid, string(body[2:])}
	return
}

func writeFull(writer io.Writer, buf []byte) error {
	total := len(buf)
	for pos := 0; pos < total; {
		n, err := writer.Write(buf[pos:])
		if err != nil {
			return err
		}
		pos += n
	}
	return nil
}

func (self *tcpSession) sendMessage(msg *tyszj.BasicMessage) {
	writer, ok := self.Raw().(io.Writer)
	if !ok || nil == writer {
		return
	}
	msgData := []byte(msg.TheContent())
	msgId := msg.ID()
	pkt := make([]byte, 2+2+len(msgData))
	binary.LittleEndian.PutUint16(pkt, uint16(len(msgData)+2))
	binary.LittleEndian.PutUint16(pkt[2:], uint16(msgId))
	copy(pkt[4:], msgData)
	writeFull(writer, pkt)
}

func (self *tcpSession) recvLoop() {
	for self.conn != nil {
		msg, err := self.readMessage()
		if err != nil {
			self.Close()
			self.PostEvent(&tyszj.RecvMsgEvent{self, &tyszj.SessionClosed{}})
			break
		}
		self.PostEvent(&tyszj.RecvMsgEvent{self, msg})
	}
}
func (self *tcpSession) sendLoop() {
	var writeList []*tyszj.BasicMessage
	for {
		writeList = writeList[0:0]
		exit := self.sendQueue.Pick(&writeList)
		for _, msg := range writeList {
			self.sendMessage(msg)
		}
		if exit {
			break
		}
	}
	self.cleanUp()
}

func (self *tcpSession) cleanUp() {
	self.cleanUpGuard.Lock()
	defer self.cleanUpGuard.Unlock()
	if self.conn != nil {
		self.conn.Close()
		self.conn = nil
	}
	self.exitSync.Done()
}

func (self *tcpSession) Start() {
	self.sendQueue.Reset()
	self.Peer().(ISessionManager).Add(self)
	self.exitSync.Add(2)
	go func() {
		self.exitSync.Wait()
		self.Peer().(ISessionManager).Remove(self)
	}()
	go self.recvLoop()
	go self.sendLoop()
}

func newSession(conn net.Conn, peer tyszj.IPeer) *tcpSession {
	ses := &tcpSession{
		conn:      conn,
		sendQueue: tyszj.NewPipe(),
		peer:      peer,
		PeerBundle: peer.(interface {
			GetBundle() *PeerBundle
		}).GetBundle(),
	}
	return ses
}
