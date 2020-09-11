package windows

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
)

type ExecuteCommandTask struct {
	task.Base
	Command string
}

func (t *ExecuteCommandTask) GetCmdType() definition.CmdType {
	return definition.ExecuteCommand
}

func (t *ExecuteCommandTask) GetSerializedParam() string {
	return t.Command
}
