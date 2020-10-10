package network

import (
	"fmt"
	"heartbeat-go/codec"
	"heartbeat-go/command"
	"net"
)

type HeartBeatServer struct {
	addr string
}

func RunReceiver() {
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	decoder := codec.HeartBeatUdpDecoder{}
	register := command.GetInstance()

	msg := make([]byte, 1024)
	for {
		i, retAddr, err := conn.ReadFrom(msg)
		var c *command.HeartBeatCommand
		if i < 1024 {
			c = decoder.Decode(msg[:i])
		}
		if err != nil {
			continue
		}
		handler := register.GetHandler(c.CommandName)
		fmt.Println(retAddr)
		go handler.HandleCommand(c)
	}
}
