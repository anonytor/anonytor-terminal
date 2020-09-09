package connection

import (
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type ControlConn struct {
	BaseConn
	// TokenPool manages tokens
	HostID       string
	TaskPool     task.Pool
	responseChan chan *definition.Response
	// TransConnChan 被控制连接使用，用于接收/分发传输连接
	//TransConnChan chan *TransConn
}

func (bc *BaseConn) ExpandToControlConn(hostID string) *ControlConn {
	cc := ControlConn{}
	cc.BaseConn = *bc
	cc.HostID = hostID
	return &cc
}
func (cc *ControlConn) SendTask(task task.Interface) error {
	log.Debug("trying to send cmd")
	data, _ := json.Marshal(definition.Request{
		TaskID: task.GetId(),
		Cmd:    task.GetCmdType(),
		Param:  task.GetSerializedParam(),
	})
	err := cc.basicSend(data)
	if err != nil {
		log.Warn(definition.SendError, err)
		return err
	}
	return nil
}
func (cc *ControlConn) Serve() {
	go func() {
		for {
			select {
			case <-cc.stopSignal:
				return
			case response := <-cc.responseChan:
				// 索引到task
				task, exist := cc.TaskPool.Get(response.TaskID)
				if !exist {
					continue
				}
				switch response.TaskStatus {
				case definition.TaskReceived:
					task.OnTaskReceived()
				case definition.TaskFinished:
					task.OnTaskFinished()
				case definition.TaskWantRetrieveThroughCtrl:
					task.OnTaskWantRetrieveThroughCtrl(response.Data)
				case definition.TaskWantRetrieveThroughTrans:
					task.OnTaskWantRetrieveThroughTrans()
				}
			}
		}

	}()

}

// ListenIncome should run in a goroutine
// which recv continuously from connection
// and tries its best to UnMarshal the byte
// if UnMarshal succeed, it sends the pointer of
// type needed through chan
func (cc *ControlConn) ListenIncome() {
	for {
		select {
		case <-cc.stopSignal:
			close(cc.stopSignal)
			return
		default:
			b, err := cc.basicRecv()
			if err != nil {
				log.Warn(err)
				continue
			}
			// 反序列化
			r := definition.Response{}
			err = json.Unmarshal(b, &r)
			if err != nil {
				log.Warn(err)
				continue
			}
			// 压入 chan 中
			cc.responseChan <- &r
		}
	}
}
