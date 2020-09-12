package connection

import (
	"anonytor-terminal/runtime/definition"
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	MaxBuffLen     = 4096
	IncreaseRate   = 1.5
	DefaultBuffLen = 128
	ChanBuffSize   = 64
)
const (
	TransTimeOut     = 100 * time.Second
	HandshakeTimeOut = TransTimeOut
)

type BaseConn struct {
	db     *gorm.DB
	addr   string
	buffer []byte
	//PreparedBuffer  []byte
	Conn          net.Conn
	mutex         sync.Mutex
	handshakeChan chan *definition.Handshake
	//DecodeErrChan   signal
	stopSignal definition.Signal
}

func NewBaseConn(conn net.Conn, db *gorm.DB) *BaseConn {
	bc := BaseConn{}
	bc.buffer = make([]byte, 0, DefaultBuffLen)
	bc.Conn = conn
	bc.db = db
	bc.addr = conn.RemoteAddr().String()
	bc.handshakeChan = make(chan *definition.Handshake, ChanBuffSize)
	bc.stopSignal = make(definition.Signal, ChanBuffSize)
	return &bc
}

func (bc *BaseConn) Close() {
	// 仅发送信号，不直接关闭通道
	bc.stopSignal <- struct{}{}
	// 关闭连接
	_ = bc.Conn.Close()
}
func (bc *BaseConn) Serve() {
	go func() {
		b, err := bc.basicRecv()
		if err != nil {
			log.Warn(err)
		}
		hs := definition.Handshake{}
		err = json.Unmarshal(b, &hs)
		if err != nil {
			log.Warn(err)
		}
		bc.handshakeChan <- &hs
		return
	}()
}
func (bc *BaseConn) RecvHandshake() (*definition.Handshake, error) {
	timer := time.NewTimer(HandshakeTimeOut)
	select {
	case <-timer.C:
		return nil, definition.TimeOutError
	case hs := <-bc.handshakeChan:
		if hs.HostID != "" {
			return hs, nil
		} else {
			return nil, definition.TimeOutError
		}
	}
}
func (bc *BaseConn) OK() error {
	c := definition.Request{}
	c.Cmd = definition.OK
	content, _ := json.Marshal(c)
	err := bc.basicSend(content)
	return err
}

// Reset tries its best to send ResetCmd to client
// and the close the connect
func (bc *BaseConn) Reset() {
	c := definition.Request{}
	c.Cmd = definition.Reset
	content, _ := json.Marshal(c)
	_ = bc.basicSend(content)
	_ = bc.Conn.Close()
}

/////// STATIC METHODS
/////// STATIC METHODS
/////// STATIC METHODS
// basicRecv recv bytes from connection and
// returns the slice of received bytes
func (bc *BaseConn) basicRecv() ([]byte, error) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()
	// 从缓冲区读入buffer
	_, err := bc.readAndAppendToBuffer()
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	// \n 所处的位置
	index := bytes.Index(bc.buffer, []byte("\n"))
	if index == -1 {
		log.Warn(definition.DataFormatError)
		return nil, definition.DataFormatError
	}
	// 获取真正 data 的长度 dataLen
	dataLen, err := strconv.Atoi(string(bc.buffer[0:index]))
	if err != nil {
		log.Warn(err)
		return nil, definition.DataFormatError
	} else if dataLen <= 1 {
		// data长度不合法
		log.Warn(definition.PayloadSizeError, dataLen)
		return nil, definition.DataFormatError
	}
	// 此时可对预期的边界进行计算
	expectLen := dataLen + index + 1
	// 不符合预期时，每次将新读入的数据追加到 pc.buffer 的末尾
	for len(bc.buffer) < expectLen {
		_, err := bc.readAndAppendToBuffer()
		if err != nil {
			log.Warn(err)
		}
	}
	// 正确
	log.Debug("complete read from system buffer")
	// 真正的 data，为 \n 后面到结尾处
	result := bc.buffer[index+1 : expectLen]
	// 此时由于可能发生粘包的情况，后面的内容不能被丢弃
	tailBuf := bc.buffer[expectLen:]
	// 根据 tailBuf 大小为Buffer分配新的大小
	// 将过长的 buffer 进行垃圾回收（自动）
	// 但保证 buffer 的大小大于 DefaultBuffLen
	if len(tailBuf) == 0 {
		bc.refreshBuffer()
	} else {
		bc.buffer = tailBuf
	}
	return result, nil
}
func (bc *BaseConn) basicSend(data []byte) error {
	paddedData := bc.padData(data)
	_, err := bc.Conn.Write(paddedData)
	return err
}
func (bc *BaseConn) padData(raw []byte) []byte {
	rawLen := len(raw)
	buf := make([]byte, 0, rawLen)
	buf = append(buf, []byte(strconv.Itoa(rawLen))...)
	buf = append(buf, '\n')
	buf = append(buf, raw[:rawLen]...)
	return buf
}
func (bc *BaseConn) readAndAppendToBuffer() (int, error) {
	// tmpBufCap 最大不超过4096，防止爆栈
	tmpBufCap := 0
	crtBufCap := cap(bc.buffer)
	if crtBufCap > MaxBuffLen {
		tmpBufCap = MaxBuffLen
	} else {
		// 在合理的范围内，每次为上次的 IncreaseRate 倍
		tmpBufCap = int(float64(crtBufCap) * IncreaseRate)
	}
	tmpBuffer := make([]byte, tmpBufCap)
	_ = bc.Conn.SetReadDeadline(time.Now().Add(TransTimeOut))
	n, err := bc.Conn.Read(tmpBuffer)
	if err != nil {
		return 0, err
	}
	// 仅将切片的可用的部分加入 pc.buffer
	bc.buffer = append(bc.buffer, tmpBuffer[:n]...)
	return n, nil
}
func (bc *BaseConn) refreshBuffer() {
	bc.buffer = make([]byte, 0, DefaultBuffLen)
}
