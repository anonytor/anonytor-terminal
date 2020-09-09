package control

import (
	"anonytor-terminal/control/windows"
	"anonytor-terminal/runtime/definition"
	"github.com/imdario/mergo"
	log "github.com/sirupsen/logrus"
	"net"
)

var (
	handlers definition.HandlerMap
)

func registerHandlers() {
	_ = mergo.Map(&handlers, windows.Handlers)
}

func handleConnection(conn net.Conn) {
	//退出后关闭连接
	defer conn.Close()
	//1. 接收握手数据
	payload, err := recvHandshakeFromConn(conn)
	if err != nil {
		log.Warn("can't finish payload with client: ", err)
		return
	}
	err = sendJsonToConn(conn, definition.H{
		"status": "ok",
	})
	if err != nil {
		log.Warn(definition.SendError)
	}
	handler, exist := handlers[payload.Cmd]
	if !exist {
		return
	}
	handler()
}

//func recvFile()
