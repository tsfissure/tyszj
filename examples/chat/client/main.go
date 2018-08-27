package main

import (
	"fmt"
	"time"
	"tyszj"
	"tyszj/examples/chat/msg/pbgo"

	_ "tyszj/peer"

	"github.com/gogo/protobuf/proto"
)

func onMessage(ev tyszj.IEvent) {

}

func readConsole(peer tyszj.IPeer) {
	msg := &pbgo.FirstProto{
		Value: 2333,
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("Marshal error:", msg)
		return
	}
	peer.(interface {
		Session() tyszj.ISession
	}).Session().Send(tyszj.BasicMessage{
		Id:      1001,
		Content: string(data),
	})
}

func main() {
	queue := tyszj.NewEventQueue()

	peer := tyszj.NewGenericPeer("tcp.Connector", "127.0.0.1:2333", queue, onMessage)

	peer.Start()

	queue.StartLoop()
	time.Sleep(2 * time.Second)
	readConsole(peer)
	queue.Wait()
}
