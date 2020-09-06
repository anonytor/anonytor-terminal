package control

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net"
)

type Socket struct {
	db *gorm.DB
	bindAddr string
}

func InitSocket(db *gorm.DB, addr string) *Socket {
	registerHandlers()
	return &Socket{db, addr}
}
func (socket *Socket) ListenAndServe() {
	listener, err := net.Listen("tcp", socket.bindAddr)
	if err != nil {
		log.Fatal(
			fmt.Sprintf("can't start tcp server, because %v", err,
			))

	} else {
		log.Debug("tcp server started, listening on ", socket.bindAddr)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Warn(
				fmt.Sprintf("control error: %v", err),
			)
		} else {
			log.Debug(
				fmt.Sprintf("connection established from %s", conn.RemoteAddr()),
			)
		}
		go handleConnection(conn)
	}
}
