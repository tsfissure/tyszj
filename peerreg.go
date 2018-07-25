package tyszj

type FPeerCreateFunc func() IPeer

var gCreatorByTypeName = map[string]FPeerCreateFunc{}

func RegisterPeerCreator(f FPeerCreateFunc) {
	tempPeer := f()
	if _, ok := gCreatorByTypeName[tempPeer.TypeName()]; ok {
		panic("Dumplicate peer type" + tempPeer.TypeName())
	}

	gCreatorByTypeName[tempPeer.TypeName()] = f
}

func NewPeer(peerType string) IPeer {
	creator, ok := gCreatorByTypeName[peerType]
	if !ok {
		panic("peerType " + peerType + " not exists")
	}
	return creator()
}

func NewGenericPeer(peerType, addr string, q IEventQueue, f FEventCallback) GenericPeer {
	p := NewPeer(peerType)
	gp := p.(GenericPeer)
	gp.SetAddress(addr)
	gp.SetQueue(q)
	bundle := p.(IProcessorBundle)
	bundle.SetCallback(f)
	return gp
}
