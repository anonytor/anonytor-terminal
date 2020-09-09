package connection

import "anonytor-terminal/controller/task"

type TransConn struct {
	// TaskID 仅在封装传输连接时使用
	Task task.Interface
	BaseConn
}

func (tc *TransConn) ListenIncome() {

}
func (tc *TransConn) Serve() {
	go func() {
		tc.Task.OnTransConnEstablished(tc.Conn)
	}()
}
func (bc *BaseConn) ExpandToTransferConn(task task.Interface) *TransConn {
	tc := TransConn{}
	tc.BaseConn = *bc
	tc.Task = task
	return &tc
}
