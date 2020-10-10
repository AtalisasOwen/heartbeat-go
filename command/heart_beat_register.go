package command

import (
	"sync"
)

type HeartBeatCommandRegister struct {
	handlers map[string]HeartBeatCommandHandler
	mu       sync.Mutex
}

var instance *HeartBeatCommandRegister
var once sync.Once

func GetInstance() *HeartBeatCommandRegister {
	once.Do(func() {
		instance = &HeartBeatCommandRegister{
			handlers: make(map[string]HeartBeatCommandHandler),
		}
	})

	return instance
}

func (register *HeartBeatCommandRegister) AddHandler(commandName string, handler HeartBeatCommandHandler) {
	register.mu.Lock()
	register.handlers[commandName] = handler
	register.mu.Unlock()
}

func (register *HeartBeatCommandRegister) GetHandler(commandName string) HeartBeatCommandHandler {
	register.mu.Lock()
	defer register.mu.Unlock()
	handler, ok := register.handlers[commandName]
	if !ok {
		panic("Don't find CommandHandler for Command: " + commandName)
	}
	return handler
}
