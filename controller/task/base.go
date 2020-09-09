package task

import (
	"anonytor-terminal/runtime/definition"
	"io"
)

type Interface interface {
	GetCmdType() definition.CmdType
	GetId() string
	SetId(string)
	GetSerializedParam() string
	// Lifetimes
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
}

func (bt *Base) GetId() string {
	return bt.ID
}

func (bt *Base) SetId(id string) {
	bt.ID = id
}

func (bt *Base) OnTaskReceived() {
	bt.Status = definition.TaskReceived
}

//func (bt *Base)
