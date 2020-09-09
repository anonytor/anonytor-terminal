package handlers

import (
	"anonytor-terminal/api/middlewares"
	"anonytor-terminal/database/models"
	"anonytor-terminal/runtime/definition"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterInfo(r *gin.RouterGroup) {
	r.GET("", GetInfo())
}

func GetInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := middlewares.GetDb(c)
		var (
			hostCount       uint
			connectionCount uint
		)
		if v := db.Model(&models.Host{}).Count(&hostCount); v.Error != nil {
			panic(v.Error)
		}
		if v := db.Model(&models.Connection{}).Count(&connectionCount); v.Error != nil {
			panic(v.Error)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":          definition.StatusOK,
			"hostCount":       hostCount,
			"connectionCount": connectionCount,
		})
	}
}
