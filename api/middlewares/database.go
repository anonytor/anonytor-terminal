package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetDb(c *gin.Context) *gorm.DB {
	return c.MustGet("db").(*gorm.DB)
}

func SetDb(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
