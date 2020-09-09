package definition

import (
	"net"
)

type HandleFunc func(conn net.Conn)

type HandlersMap map[CmdType]HandleFunc
type H map[string]interface{}
