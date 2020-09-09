package handlers

import (
	"github.com/gin-gonic/gin"
	"monitor-server-backend/api/middlewares"
	"monitor-server-backend/database/models"
	"net/http"
)

func GetConnectionList() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := middlewares.GetDb(c)
		id := c.Param("id")
		connections := models.GetConnectionsByHostId(db, id)
		c.JSON(http.StatusOK, gin.H{
			"status":   0,
			"resource": connections,
		})
	}
}
