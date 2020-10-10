package command

import (
	"fmt"
	"net"
)

const SEPARATOR = ":"

type HeartBeatCommandHandler interface {
	HandleCommand(command *HeartBeatCommand)
}

type HeartBeatCommand struct {
	SrcAddr     net.Addr
	DstAddr     net.Addr
	CommandName string
	CommandUuid string
	Data        []byte
}

func (a *HeartBeatCommand) String() string {
	if a == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HeartBeatCommand{SrcAddr=%s, DstAddr=%s, commandName=%s, commandUuid=%s}",
		a.SrcAddr, a.DstAddr, a.CommandName, a.CommandUuid)
}
