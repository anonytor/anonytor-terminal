package handlers

import (
	"anonytor-terminal/runtime/definition"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAgent(r *gin.RouterGroup) {
	r.GET("", DownloadAgent())
}

func DownloadAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		os := c.Param("os")
		id := c.Param("id")
		fmt.Println(os, id)
		c.JSON(http.StatusNotFound, gin.H{
			"status": definition.StatusNotFound,
		})
	}
}
