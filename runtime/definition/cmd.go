package definition

type CmdType int

const (
	UploadFile = iota
	GetFileContent
	KeyboardInputRecordUpload
	TestUpload
	Confirm
	Reset
)
