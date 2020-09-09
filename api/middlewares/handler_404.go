package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler404() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": -1,
		})
	}
}
