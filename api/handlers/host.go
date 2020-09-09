package handlers

import (
	"anonytor-terminal/api/middlewares"
	"anonytor-terminal/database/models"
	"anonytor-terminal/runtime/definition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHost(r *gin.RouterGroup) {
	r.GET("", GetHostList())
	r.GET(":id", GetHostDetail())
	r.POST("", CreateHost())
	r.DELETE(":id", DeleteHost())
}

func GetHostList() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := middlewares.GetDb(c)
		var hosts []models.Host
		if v := db.Order("created_at desc").Find(&hosts); v.Error != nil {
			panic(v.Error)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
			"hosts":  hosts,
		})
	}
}

func GetHostDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := middlewares.GetDb(c)
		id := c.Param("id")
		host := models.GetHostById(db, id)
		if host == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": definition.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
			"host":   host,
		})
	}
}

func CreateHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Name string `json:"name"`
		}
		var r req
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
		}
		db := middlewares.GetDb(c)
		host := models.NewHost(db, r.Name)
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
			"id":     host.ID,
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
				"status": definition.StatusNotFound,
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
			"status": definition.StatusOK,
		})
	}
}
