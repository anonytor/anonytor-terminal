package definition

import "errors"

var (
	SendError           = errors.New("can't send through connection")
	CmdNotExistError    = errors.New("such cmd not provided")
	CmdNotReceivedError = errors.New("can't receive cmd from connection")
	DataFormatError     = errors.New("data format error")
	PayloadSizeError    = errors.New("wrong payload size")
	TimeOutError        = errors.New("timed out ")
	NoSuchConnError     = errors.New("no such control connection in pool")
)
