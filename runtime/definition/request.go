package definition

type Request struct {
	TaskID string `json:"task_id"`
	Cmd CmdType `json:"cmd"`
	Param string `json:"param"`
}