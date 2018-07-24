package tyszj

type IEvent interface {
	//事件对应的session
	Session() ISession

	//事件内容
	Message() interface{}
}

type FEventCallback func(ev IEvent)
