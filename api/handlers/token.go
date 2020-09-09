package handlers

import (
	"github.com/gin-gonic/gin"
	"monitor-server-backend/api/middlewares"
	"monitor-server-backend/database/models"
	"net/http"
	"time"
)

func CreateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		type req struct {
			ExpiredAt time.Time `json:"expired_at"`
		}
		var r req
		err := c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
			})
			return
		}
		if r.ExpiredAt.Before(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -2,
			})
			return
		}
		db := middlewares.GetDb(c)
		token := models.NewToken(db, r.ExpiredAt)
		c.JSON(http.StatusOK, gin.H{
			"status":     0,
			"token":      token,
			"expired_at": token.ExpireAt,
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
				"status": -1,
			})
			return
		}
		db := middlewares.GetDb(c)
		models.DeleteToken(db, r.Token)
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
		})
	}
}
