package windows

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
)

type GetClipboardTextTask struct {
	task.Base
	Result chan string
}

func (t *GetClipboardTextTask) GetCmdType() definition.CmdType {
	return definition.GetClipboardText
}

func (t *GetClipboardTextTask) OnTaskWantRetrieveThroughCtrl(buf []byte) {
	t.Result <- string(buf)
	t.GetPool().Remove(t)
}
