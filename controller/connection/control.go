package connection

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
)

type ControlConn struct {
	BaseConn
	// TokenPool manages tokens
	HostID           string
	TaskPool         task.Pool
	responseChan     chan *definition.Response
	reportBrokenChan chan string
	// TransConnChan 被控制连接使用，用于接收/分发传输连接
	//TransConnChan chan *TransConn
}

func (bc *BaseConn) ExpandToControlConn(hostID string, rbc chan string) *ControlConn {
	cc := ControlConn{}
	cc.BaseConn = *bc
	cc.HostID = hostID
	cc.reportBrokenChan = rbc
	return &cc
}
func (cc *ControlConn) SendTask(task task.Interface) error {
	log.Debug("trying to add task to control connection's taskPool")
	cc.TaskPool.Add(task)
	log.Debug("task added to pool")
	log.Debug("trying to send cmd")
	r := definition.Request{
		TaskID: task.GetId(),
		Cmd:    task.GetCmdType(),
		Param:  task.GetSerializedParam(),
	}
	err := cc.SendRequest(r)
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
			default:
				b, err := cc.basicRecv()
				if err != nil {
					if err != definition.TimeOutError {
						log.Warn("control connection is totally broken, exit")
						cc.stopSignal <- struct{}{}
						cc.reportBrokenChan <- cc.HostID
						break
					}
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
				if r.TaskID == "" {
					continue
				} else {
					_ = cc.OK()
				}
				// 索引到task
				t, exist := cc.TaskPool.Get(r.TaskID)
				if !exist {
					continue
				}
				switch r.TaskStatus {
				case definition.TaskReceived:
					t.OnTaskReceived()
				case definition.TaskFinished:
					t.OnTaskFinished(r.Data)
				case definition.TaskWantRetrieveThroughTrans:
					t.OnTaskWantRetrieveThroughTrans()
				}
			}
		}

	}()

}
func (cc *ControlConn) SendRequest(request definition.Request) error {
	data, _ := json.Marshal(request)
	err := cc.basicSend(data)
	return err
}
