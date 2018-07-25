package main

import (
	"tyszj"
	_ "tyszj/peer"
)

func onMessage(ev tyszj.IEvent) {

}

func main() {

	queue := tyszj.NewEventQueue()

	p := tyszj.NewGenericPeer("tcp.Acceptor", "127.0.0.1:2333", queue, onMessage)
	p.Start()
	queue.StartLoop()
	queue.Wait()
}
