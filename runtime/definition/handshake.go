package definition

type ConnType int

const (
	ControlConn = iota
	TransferConn
)

type Handshake struct {
	HostID string `json:"host_id"`
	Key    string   `json:"key"`
	Type   ConnType `json:"type"`
	TaskID string   `json:"task_id"`
}
