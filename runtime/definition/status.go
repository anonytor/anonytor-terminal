package definition

const (
	StatusOK = iota + 10000
	StatusNotFound
	StatusServerError
	StatusPermissionDenied
	StatusBadRequest
	StatusExpiredAtBeforeNow
)
