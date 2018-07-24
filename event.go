package tyszj

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
