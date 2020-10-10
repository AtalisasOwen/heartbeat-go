package command

import (
	"fmt"
)

type PingCommandHandler struct{}

func (handler *PingCommandHandler) HandleCommand(command *HeartBeatCommand) {
	fmt.Printf("Getting Ping from %s\n", command)
}
