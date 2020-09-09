package definition

type TaskStatus int

const (
	TaskInitialized = iota
	// 发送前
	TaskSent
	// 发送后
	TaskReceived
	// 二选一，是否提升连接
	TaskWantRetrieveThroughCtrl
	TaskWantRetrieveThroughTrans
	// 传输连接建立
	TransConnEstablished
	// 完成
	TaskFinished

	TaskErrorInExecution
)

type TaskInfo struct {
	//TaskID  string     `json:"token"`
	Status  TaskStatus `json:"status"`
	Content []byte     `json:"content"`
}
