package tyszj

type BasicMessage struct {
	Id      uint16
	Content string
}

func (self *BasicMessage) ID() uint16 {
	return self.Id
}
func (self *BasicMessage) TheContent() string {
	return self.Content
}

type RecvMsgEvent struct {
	Ses ISession
	Msg interface{}
}

func (self *RecvMsgEvent) Session() ISession {
	return self.Ses
}

func (self *RecvMsgEvent) Message() interface{} {
	return self.Msg
}

type SendMsgEvent struct {
	Ses ISession
	Msg interface{}
}

func (self *SendMsgEvent) Session() ISession {
	return self.Ses
}

func (self *SendMsgEvent) Message() interface{} {
	return self.Msg
}
