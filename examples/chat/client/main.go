package main

import (
	"tyszj"
)

func onMessage(ev tyszj.IEvent) {

}

func readConsole() {
	for {

	}
}

func main() {
	queue := tyszj.NewEventQueue()

	peer := tyszj.NewGenericPeer("tcp.Connector", "127.0.0.1:2333", queue, onMessage)

	peer.Start()

	queue.StartLoop()
	readConsole()
}
