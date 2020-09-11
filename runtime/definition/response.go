package definition

type Response struct {
	TaskID     string     `json:"task_id"`
	TaskStatus TaskStatus `json:"task_status"`
	Data       string     `json:"data"`
}
type TaskStatus int

const (
	TaskInitialized = iota
	// 发送前
	TaskSent
	// 发送后
	TaskReceived
	// 二选一，是否提升连接
	TaskWantRetrieveThroughTrans
	// 传输连接建立
	TaskTransConnEstablished
	// 完成
	TaskFinished

	TaskErrorInExecution
)