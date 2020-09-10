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
