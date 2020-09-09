package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLoggedIn := session.Get("isLoggedIn").(bool)
		if isLoggedIn {
			expiredAt := session.Get("expiredAt").(time.Time)
			if time.Now().After(expiredAt) {
				session.Set("isLoggedIn", false)
				_ = session.Save()
			}
		}
		c.Next()
	}
}

func RequireLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLoggedIn := session.Get("isLoggedIn").(bool)
		if !isLoggedIn {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": -1,
			})
		} else {
			c.Next()
		}
	}
}
