package handlers

import (
	"github.com/gin-gonic/gin"
	"monitor-server-backend/api/middlewares"
	"monitor-server-backend/database/models"
	"net/http"
)

func CreateHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Name string
		}
		var r req
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
			})
		}
		db := middlewares.GetDb(c)
		host := models.NewHost(db, r.Name)
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"id":     host.ID,
			"key":    host.Key,
		})
	}
}

func DeleteHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		tx := middlewares.GetDb(c).Begin()
		defer tx.RollbackUnlessCommitted()
		host := models.GetHostById(tx, id)
		if host == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": -1,
			})
			return
		}
		if v := tx.Delete(host); v.Error != nil {
			panic(v.Error)
		}
		if v := tx.Commit(); v.Error != nil {
			panic(v.Error)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
		})
	}
}
