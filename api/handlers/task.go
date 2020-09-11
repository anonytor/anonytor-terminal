package handlers

import (
	"anonytor-terminal/api/middlewares"
	"anonytor-terminal/controller/task/windows"
	"anonytor-terminal/runtime/definition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterTask(r *gin.RouterGroup) {
	r.POST("", CreateTask())
	r.POST("status", CheckTask())
}

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			ID      string
			CmdType definition.CmdType     `json:"cmdType"`
			Params  map[string]interface{} `json:"params"`
		}
		var r req
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
			return
		}
		ctrl := middlewares.GetController(c)
		switch r.CmdType {
		case definition.GetFileContent:
			task := windows.GetFileContentTask{Path: r.Params["path"].(string)}
			err := ctrl.ExecuteTask(r.ID, &task)
			if err != nil {
				panic(err)
			}
		}
	}
}

func CheckTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			HostId string `json:"host_id"`
			TaskID string `json:"task_id"`
		}
		var r req
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
			return
		}
		ctrl := middlewares.GetController(c)
		cc := ctrl.GetControlConnection(r.HostId)
		if cc == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
			return
		}
		task, exist := cc.TaskPool.Get(r.TaskID)
		if !exist {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
			return
		}
		switch task.GetCmdType() {
		case definition.GetClipboardText:
			clipTask := task.(ClipboardTask)
			c.JSON(http.StatusOK, gin.H{
				"status": task.getStatus,
			})
		}
	}
}
