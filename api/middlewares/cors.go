package middlewares

import (
	"anonytor-terminal/runtime/definition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status": definition.StatusOK,
			})
		}
		c.Next()
	}
}
