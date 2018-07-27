package tyszj

type BasicMessage struct {
	ID      int32
	content string
}

type RecvMsgEvent struct {
	Ses ISession
	Msg BasicMessage
}

func (self *RecvMsgEvent) Session() ISession {
	return self.Ses
}

func (self *RecvMsgEvent) Message() interface{} {
	return self.Msg
}

type SendMsgEvent struct {
	Ses ISession
	Msg BasicMessage
}

func (self *SendMsgEvent) Session() ISession {
	return self.Ses
}

func (self *SendMsgEvent) Message() interface{} {
	return self.Msg
}
