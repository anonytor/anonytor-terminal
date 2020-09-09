package handlers

import (
	"anonytor-terminal/runtime/definition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPing(r *gin.RouterGroup) {
	r.GET("", Ping())
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
		})
	}
}
