package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"anonytor-terminal/api/middlewares"
	"anonytor-terminal/database/models"
	"anonytor-terminal/runtime/definition"
)

func RegisterConnection(r *gin.RouterGroup) {
	r.GET("", GetConnectionList())
}

func GetConnectionList() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := middlewares.GetDb(c)
		id := c.Query("id")
		connections := models.GetConnectionsByHostId(db, id)
		c.JSON(http.StatusOK, gin.H{
			"status":      definition.StatusOK,
			"connections": connections,
		})
	}
}
