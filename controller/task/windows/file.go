package windows

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
	"io"
)

type GetFileContentTask struct {
	task.Base
	Path string
}

func (t *GetFileContentTask) GetCmdType() definition.CmdType {
	return definition.GetFileContent
}

func (t *GetFileContentTask) GetSerializedParam() string {
	return t.Path
}

func (t *GetFileContentTask) OnTransConnEstablished(r io.ReadWriter) {
	// TODO: 打开文件流
}

type UploadFileTask struct {
	task.Base
	targetPath string
	size       int
}

func NewUploadFileTask(path string, targetPath string) {
	// TODO: 读取文件信息（大小）
}

func (t *UploadFileTask) GetCmdType() definition.CmdType {
	return definition.GetFileContent
}

func (t *UploadFileTask) OnTransConnEstablished(r io.ReadWriter) {
	// TODO: 写文件
}
