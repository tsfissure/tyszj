package main

import (
	"fmt"
	"tyszj"
	"tyszj/examples/chat/msg/pbgo"
	_ "tyszj/peer"

	"github.com/gogo/protobuf/proto"
)

func onMsgInternal(ses tyszj.ISession, msg *tyszj.BasicMessage) {
	switch msg.ID() {
	case 1001:
		fp := &pbgo.FirstProto{}
		err := proto.Unmarshal([]byte(msg.TheContent()), fp)
		if err == nil {
			fmt.Println("FirstProto:", fp.GetValue())
		}
	}
}

func onMessage(ev tyszj.IEvent) {
	switch msg := ev.Message().(type) {
	case *tyszj.SessionAccepted:
		fmt.Println("accepted new session:", ev.Session().ID())
	case *tyszj.SessionClosed:
		fmt.Println("closed session", ev.Session().ID())
	case *tyszj.BasicMessage:
		onMsgInternal(ev.Session(), ev.Message().(*tyszj.BasicMessage))
	default:
		fmt.Println("Unknown Type", msg)
	}
}

func main() {

	queue := tyszj.NewEventQueue()

	p := tyszj.NewGenericPeer("tcp.Acceptor", "127.0.0.1:2333", queue, onMessage)
	p.Start()
	queue.StartLoop()
	queue.Wait()
	fmt.Println("Done...")
}
