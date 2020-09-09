package handlers

import (
	"anonytor-terminal/api/middlewares"
	"anonytor-terminal/database/models"
	"anonytor-terminal/runtime/definition"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterToken(r *gin.RouterGroup) {
	r.GET("", GetTokenList())
	r.POST("", CreateToken())
	r.DELETE("", DeleteToken())
}

func GetTokenList() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := middlewares.GetDb(c)
		var tokens []models.Token
		if v := db.Order("created_at desc").Find(&tokens); v.Error != nil {
			panic(v.Error)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
			"tokens": tokens,
		})
	}
}

func CreateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			ExpiredAt time.Time `json:"expiredAt"`
		}
		var r req
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
			return
		}
		if r.ExpiredAt.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusExpiredAtBeforeNow,
			})
			return
		}
		db := middlewares.GetDb(c)
		token := models.NewToken(db, r.ExpiredAt)
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
			"token":  token,
		})
	}
}

func DeleteToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			Token string `json:"token"`
		}
		var r req
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": definition.StatusBadRequest,
			})
			return
		}
		db := middlewares.GetDb(c)
		models.DeleteToken(db, r.Token)
		c.JSON(http.StatusOK, gin.H{
			"status": definition.StatusOK,
		})
	}
}
