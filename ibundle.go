package tyszj

type IMessagePoster interface {
	PostEvent(ev IEvent)
}

type IProcessorBundle interface {
	SetCallback(v FEventCallback)
}
