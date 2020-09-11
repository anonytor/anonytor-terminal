package task

import (
	"anonytor-terminal/runtime/definition"
	"io"
)

type Interface interface {
	GetId() string
	SetId(string)
	GetCmdType() definition.CmdType
	GetSerializedParam() string
	GetPool() *Pool
	SetPool(*Pool)
	// Lifetimes
	OnTaskInitialized()
	OnTaskReceived()
	OnTaskWantRetrieveThroughCtrl([]byte)
	OnTaskWantRetrieveThroughTrans()
	OnTransConnEstablished(io.ReadWriter)
	OnTaskFinished()
}

type Base struct {
	Interface
	ID     string
	Status definition.TaskStatus
	Pool   *Pool
}

func (bt *Base) GetId() string {
	return bt.ID
}

func (bt *Base) SetId(id string) {
	bt.ID = id
}

func (bt *Base) GetPool() *Pool {
	return bt.Pool
}

func (bt *Base) SetPool(p *Pool) {
	bt.Pool = p
}

func (bt *Base) OnTaskInitialized() {
	bt.Status = definition.TaskWantRetrieveThroughCtrl
}

func (bt *Base) GetSerializedParam() string {
	return ""
}

//func (bt *Base)
