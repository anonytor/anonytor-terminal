package middlewares

import (
	"anonytor-terminal/database/models"
	"anonytor-terminal/runtime/definition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := GetDb(c)
		token := c.GetHeader("Authorization")
		if token == "" || !models.CheckToken(db, token) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": definition.StatusPermissionDenied,
			})
		}
		c.Next()
	}
}
