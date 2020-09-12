package connection

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"anonytor-terminal/controller/task"
	"anonytor-terminal/database/models"
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
						var cct models.Connection
						if v := cc.db.Where("addr = ? and host_id = ?", cc.addr, cc.HostID).First(&cct); v.Error == nil {
							if v := cc.db.Delete(&cct); v.Error != nil {
								panic(v.Error)
							}
						} else if !gorm.IsRecordNotFoundError(v.Error) {
							panic(v.Error)
						}
						var count int
						if v := cc.db.Model(&models.Connection{}).Where("host_id = ? and type = 0", cc.HostID).Count(&count); v.Error != nil {
							panic(v.Error)
						}
						if count == 0 {
							host := models.GetHostById(cc.db, cc.HostID)
							host.Status = 0
							if v := cc.db.Save(host); v.Error != nil {
								panic(v.Error)
							}
						}
						log.Warn("control connection is totally broken, exit")
						cc.stopSignal <- struct{}{}
						cc.reportBrokenChan <- cc.HostID
						break
					}
					continue
				}
				host := models.GetHostById(cc.db, cc.HostID)
				host.LastSeen = time.Now()
				if v := cc.db.Save(host); v.Error != nil {
					panic(v.Error)
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
