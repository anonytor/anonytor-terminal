package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"monitor-server-backend/api/middlewares"
	"monitor-server-backend/database/models"
	"net/http"
	"time"
)

func SetAuthHandlers(r *gin.RouterGroup) {
	r.POST("", Auth())
	r.DELETE("", middlewares.RequireLoggedIn(), Logout())
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := middlewares.GetDb(c)
		session := sessions.Default(c)
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
		if !models.CheckToken(db, r.Token) {
			// TODO: Set error code.
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
			})
		} else {
			session.Set("isLoggedIn", true)
			session.Set("expiredAt", time.Now().Add(1*time.Hour))
			_ = session.Save()
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
			})
		}
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("isLoggedIn", false)
		_ = session.Save()
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
		})
	}
}
