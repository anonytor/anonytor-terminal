package connection

import (
	"anonytor-terminal/runtime/definition"
	"sync"
)

type CtrlConnPool struct {
	ctrlConnMap sync.Map
	// 用于关闭 pool
	closeSignal definition.Signal
	// ControlConn 通过此通道向上层通报自己已损坏
	CtrlConnBrokenSignal chan string
}

func NewCtrlPool() CtrlConnPool {
	ccp := CtrlConnPool{}
	ccp.ctrlConnMap = sync.Map{}
	ccp.CtrlConnBrokenSignal = make(chan string, DefaultBuffLen)
	return ccp
}
func (ccp *CtrlConnPool) Add(pc *ControlConn) {
	ccp.ctrlConnMap.Store(pc.HostID, pc)
}
func (ccp *CtrlConnPool) Get(id string) (*ControlConn, bool) {
	cc, exist := ccp.ctrlConnMap.Load(id)
	if !exist {
		return nil, false
	}
	return cc.(*ControlConn), true

}
func (ccp *CtrlConnPool) CloseConn(uuid string) {
	ccp.ctrlConnMap.Delete(uuid)

}

// Close closes the whole ctrlConnMap
func (ccp *CtrlConnPool) Close() {
	ccp.ctrlConnMap.Range(func(key, value interface{}) bool {
		cc := value.(*ControlConn)
		cc.Close()
		return true
	})
}
func (ccp *CtrlConnPool) ListenToBroken() {
	go func() {
		for {
			select {
			case <-ccp.closeSignal:
				return
			case HostID := <-ccp.CtrlConnBrokenSignal:
				ccp.CloseConn(HostID)
			}
		}
	}()
}
func (ccp *CtrlConnPool) GetConnections() []*ControlConn {
	result := make([]*ControlConn, DefaultBuffLen)
	ccp.ctrlConnMap.Range(func(key, value interface{}) bool {
		result = append(result, value.(*ControlConn))
		return true
	})
	return result
}
