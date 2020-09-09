package controller

import (
	"anonytor-terminal/controller/task/windows"
	"anonytor-terminal/runtime/definition"
	"github.com/imdario/mergo"
)

var (
	Handlers = definition.HandlersMap{}
)

func registerHandlers() {
	_ = mergo.Map(&Handlers, windows.Handlers)
}



//func recvFile()
