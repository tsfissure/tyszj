package tyszj

type ISession interface {
	//原始socket连接
	Raw() interface{}

	//Session的归属Peer
	Peer() IPeer

	Send(msg BasicMessage)

	ID() int64
	SetID(id int64)

	Close()
}
