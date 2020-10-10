package test

import (
	"fmt"
	"heartbeat-go/codec"
	"heartbeat-go/command"
	"net"
	"testing"
)

func TestEncodeAndDecoder(t *testing.T) {
	srcAddr, _ := net.ResolveUDPAddr("udp", "localhost:6666")
	dstAddr, _ := net.ResolveUDPAddr("udp", "localhost:8888")

	c := &command.HeartBeatCommand{
		SrcAddr:     srcAddr,
		DstAddr:     dstAddr,
		CommandName: "PING",
		CommandUuid: "adasdasadasfas",
	}

	encoder := codec.HeartBeatUdpEncoder{}
	decoder := codec.HeartBeatUdpDecoder{}
	b := encoder.Encode(c)
	decodeC := decoder.Decode(b)

	fmt.Println(decodeC)
	fmt.Println(decodeC.DstAddr)
}
