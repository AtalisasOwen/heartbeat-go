package network

import (
	"fmt"
	"heartbeat-go/codec"
	"heartbeat-go/command"
	"net"
)

func RunClient() {
	fmt.Printf("client for server url: %s\n", addr)
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	msg := make([]byte, 512)
	n, err := conn.Write([]byte("connected"))
	if err != nil {
		panic(err)
	}
	for {
		n, err = conn.Read(msg)
		if err != nil {
			continue
		}
		fmt.Printf("Client Get: %s\n", string(msg[:n]))
	}
}

func RunPingClient() {
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	encoder := codec.HeartBeatUdpEncoder{}
	srcAddr, _ := net.ResolveUDPAddr("udp", "localhost:6667")
	dstAddr, _ := net.ResolveUDPAddr("udp", "localhost:8890")
	c := &command.HeartBeatCommand{
		SrcAddr:     srcAddr,
		DstAddr:     dstAddr,
		CommandName: "PING",
		CommandUuid: "adasdasadasfas",
	}
	msg := encoder.Encode(c)
	n, err := conn.Write(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
