package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// TODO: Log error
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status": -1,
				})
				return
			}
		}()
		c.Next()
	}
}
