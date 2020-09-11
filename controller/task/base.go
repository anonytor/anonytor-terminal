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
	Cmd definition.CmdType
	Param string
	Status definition.TaskStatus
}

func (bt *Base) GetId() string {
	return bt.ID
}

func (bt *Base) SetId(id string) {
	bt.ID = id
}

func (bt *Base)GetCmdType() definition.CmdType{
	return bt.Cmd
}
func (bt*Base)GetSerializedParam() string{
	return bt.Param
}
func (bt *Base) OnTaskReceived() {
	bt.Status = definition.TaskReceived
}
func (bt *Base)  OnTaskWantRetrieveThroughCtrl([]byte){
	bt.Status=definition.TaskWantRetrieveThroughCtrl
}
func (bt *Base)  OnTaskWantRetrieveThroughTrans(){
	bt.Status=definition.TaskWantRetrieveThroughTrans
}
func (bt *Base) OnTransConnEstablished(io.ReadWriter){
	bt.Status=definition.TaskTransConnEstablished
}
func (bt *Base) OnTaskFinished(){
	bt.Status=definition.TaskFinished
}