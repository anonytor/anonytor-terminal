package controller

import (
	"anonytor-terminal/controller/connection"
	"anonytor-terminal/controller/task"
	"anonytor-terminal/runtime/definition"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net"
)

type Controller struct {
	db           *gorm.DB
	bindAddr     string
	ctrlConnPool connection.CtrlConnPool
	// TODO: 增加关闭 Controller 的功能
	closeSignal definition.Signal
	TaskPool    *task.Pool
}

func InitController(db *gorm.DB, addr string) *Controller {
	registerHandlers()
	c := &Controller{}
	c.db = db
	c.bindAddr = addr
	c.ctrlConnPool = connection.NewCtrlPool()
	c.TaskPool = task.NewPool()
	return c
}

func (c *Controller) ListenAndServe() {
	// Step1：监听相应端口
	listener, err := net.Listen("tcp", c.bindAddr)
	if err != nil {
		log.Fatal(
			fmt.Sprintf("can't start tcp server, because %v", err,
			))
	} else {
		log.Debug("tcp server started, listening on ", c.bindAddr)
	}
	defer listener.Close()
	// 进入死循环
	for {
		select {
		// 可被关闭
		case <-c.closeSignal:
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Warn(
					fmt.Sprintf("c error: %v", err),
				)
			} else {
				log.Debug(
					fmt.Sprintf("connection established from %s", conn.RemoteAddr()),
				)
			}
			go c.handleConnection(conn)
		}

	}
}
func (c *Controller) handleConnection(conn net.Conn) {
	//1. 接收握手数据
	// 		实例化一个临时的 connection
	tmpConn := connection.NewBaseConn(conn)
	tmpConn.Serve()
	// 		接收握手数据
	hs, err := tmpConn.RecvHandshake()
	if err != nil {
		log.Warn("can't finish hs with client: ", err)
		tmpConn.Reset()
		return
	}
	_ = tmpConn.OK()
	if hs.Type == definition.ControlConn {
		// Expand to ControlConnect
		cc := tmpConn.ExpandToControlConn(hs.HostID)
		// Add to ControlPool
		c.ctrlConnPool.Add(cc)
		cc.Serve()
	} else if hs.Type == definition.TransferConn {
		// check if its controlConnect exists
		cc, exist := c.ctrlConnPool.Get(hs.HostID)
		if !exist {
			tmpConn.Reset()
		}
		// check if its taskID exists in controlConnect's taskPool
		targetTask, exist := cc.TaskPool.Get(hs.TaskID)
		if !exist {
			tmpConn.Reset()
		}
		// targetTask 存在，expandToTransConn
		tc := tmpConn.ExpandToTransferConn(targetTask)
		//开始后续的执行
		tc.Serve()

	}
}

//func (c *Controller) AddTransConn(uuid string, connection net.Conn) {
//	// 将conn封装
//	pc := conn_pack.baseConnPack{}.New(uuid, connection, Handlers)
//	// 加入 c 实例的连接池中
//	c.ctrlConnPool.AddConnPack(pc)
//	// 运行服务
//	go pc.Serve()
//}
func (c *Controller) ExecuteTask(id string, task task.Interface) error {
	cc, exist := c.ctrlConnPool.Get(id)
	if !exist {
		return definition.NoSuchConnError
	}
	c.TaskPool.Add(task)
	err := cc.SendTask(task)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil

}
func (c *Controller) Close() {
	c.closeSignal <- struct{}{}
	c.ctrlConnPool.Close()
}