package task

import (
	"io"

	"anonytor-terminal/runtime/definition"
)

type Interface interface {
	GetId() string
	SetId(string)
	GetStatus() definition.TaskStatus
	GetResult() string
	GetCmdType() definition.Cmd
	GetSerializedParam() string
	// Lifetimes
	OnTaskReceived()
	OnTaskWantRetrieveThroughTrans()
	OnTransConnEstablished(io.ReadWriter)
	OnTaskFinished(string)
}

type Base struct {
	Interface
	ID     string
	Cmd    definition.Cmd
	Param  string
	Status definition.TaskStatus
}

func (bt *Base) GetId() string {
	return bt.ID
}

func (bt *Base) SetId(id string) {
	bt.ID = id
}

func (bt *Base) GetStatus() definition.TaskStatus {
	return bt.Status
}

func (bt *Base) GetResult() string {
	return "Nothing here"
}

func (bt *Base) GetCmdType() definition.Cmd {
	return bt.Cmd
}
func (bt *Base) GetSerializedParam() string {
	return bt.Param
}
func (bt *Base) OnTaskReceived() {
	bt.Status = definition.TaskReceived
}
func (bt *Base) OnTaskWantRetrieveThroughTrans() {
	bt.Status = definition.TaskWantRetrieveThroughTrans
}
func (bt *Base) OnTransConnEstablished(io.ReadWriter) {
	bt.Status = definition.TaskTransConnEstablished
}
func (bt *Base) OnTaskFinished(string) {
	bt.Status = definition.TaskFinished
}
