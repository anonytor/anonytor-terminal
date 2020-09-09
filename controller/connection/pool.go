package connection

import (
	"anonytor-terminal/runtime/definition"
	"sync"
)

type CtrlConnPool struct {
	ctrlConnMap map[string]*ControlConn
	closeSignal definition.Signal
	sync.Mutex
}

func NewCtrlPool() CtrlConnPool {
	ccp := CtrlConnPool{}
	ccp.ctrlConnMap = make(map[string]*ControlConn)
	return ccp
}
func (ccp *CtrlConnPool) Add(pc *ControlConn) {
	ccp.Lock()
	{
		ccp.ctrlConnMap[pc.HostID] = pc
	}
	ccp.Unlock()
}
func (ccp *CtrlConnPool) Get(id string) (*ControlConn, bool) {
	ccp.Lock()
	defer ccp.Unlock()
	cc, exist := ccp.ctrlConnMap[id]
	return cc, exist

}
func (ccp *CtrlConnPool) CloseConn(uuid string) {
	ccp.Lock()
	{
		ccp.ctrlConnMap[uuid].Close()
		delete(ccp.ctrlConnMap, uuid)
	}
	ccp.Unlock()
}

// Close closes the whole ctrlConnMap
func (ccp *CtrlConnPool) Close() {
	ccp.Lock()
	{
		for id, cc := range ccp.ctrlConnMap {
			cc.Close()
			delete(ccp.ctrlConnMap, id)
		}
	}
	ccp.Unlock()
}
