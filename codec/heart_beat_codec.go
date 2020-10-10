package codec

import (
	"bytes"
	"heartbeat-go/command"
	"net"
	"strings"
)

type Encoder interface {
	Encode(command *command.HeartBeatCommand) []byte
}

type Decoder interface {
	Decode(buf []byte) *command.HeartBeatCommand
}

type HeartBeatUdpEncoder struct{}

type HeartBeatUdpDecoder struct{}

func (encoder *HeartBeatUdpEncoder) Encode(c *command.HeartBeatCommand) []byte {

	addr := c.SrcAddr.String()
	ipAndPort := strings.Split(addr, ":")
	buffer := bytes.Buffer{}
	buffer.WriteString(c.CommandName)
	buffer.WriteString(command.SEPARATOR)
	buffer.WriteString(c.CommandUuid)
	buffer.WriteString(command.SEPARATOR)
	buffer.WriteString(ipAndPort[0])
	buffer.WriteString(command.SEPARATOR)
	buffer.WriteString(ipAndPort[1])
	buffer.WriteString(command.SEPARATOR)
	buffer.Write(c.Data)

	return buffer.Bytes()
}

func (decoder *HeartBeatUdpDecoder) Decode(buf []byte) *command.HeartBeatCommand {
	datas := bytes.Split(buf, []byte(":"))
	commandName := datas[0]
	commandUuid := datas[1]
	ip := datas[2]
	port := datas[3]
	data := datas[4]
	addr := string(ip) + ":" + string(port)
	srcAddr, _ := net.ResolveUDPAddr("udp", addr)
	c := &command.HeartBeatCommand{
		SrcAddr:     srcAddr,
		DstAddr:     nil,
		CommandName: string(commandName),
		CommandUuid: string(commandUuid),
		Data:        data,
	}
	return c
}
