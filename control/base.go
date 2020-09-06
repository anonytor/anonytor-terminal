package control

import (
	"encoding/json"
	"fmt"
	"monitor-server-backend/runtime/definition"
	"net"
)

func recvHandshakeFromConn(conn net.Conn) (*definition.Payload, error) {
	//	1. 接收握手数据
	rawData, err := recvFromConn(conn)
	if err != nil {
		return nil, err
	}
	// 反序列化
	handshake := definition.Payload{}
	err = json.Unmarshal(rawData, &handshake)
	if err != nil {
		return nil, err
	}
	return &handshake, nil
}

func recvFromConn(conn net.Conn) ([]byte, error) {
	tmpBuff := make([]byte, 1024)
	offset, err := conn.Read(tmpBuff)
	if err != nil {
		return nil, err
	}
	dataReceived := make([]byte, offset)
	copy(dataReceived, tmpBuff)
	return dataReceived, nil
}

func sendJsonToConn(conn net.Conn, obj interface{}) error {
	buffer, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = sendToConn(conn, buffer)
	return err

}

func sendToConn(conn net.Conn, b []byte) error {
	fmt.Println(b)
	_, err := conn.Write(b)
	return err
}
