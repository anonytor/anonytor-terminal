package definition

type Response struct {
	TaskID     string     `json:"task_id"`
	TaskStatus TaskStatus `json:"task_status"`
	Data       []byte     `json:"data"`
}
