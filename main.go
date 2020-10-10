package main

import (
	"heartbeat-go/command"
	"heartbeat-go/network"
)

func main() {
	// go network.RunClient()

	register := command.GetInstance()
	var handler command.HeartBeatCommandHandler = &command.PingCommandHandler{}
	register.AddHandler("PING", handler)

	network.RunReceiver()

}
