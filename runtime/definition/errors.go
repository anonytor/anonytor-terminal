package definition

import "errors"

var (
	SendError = errors.New("can't send through conn")
)
