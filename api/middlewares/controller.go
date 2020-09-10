package middlewares

import (
	"anonytor-terminal/controller"
	"github.com/gin-gonic/gin"
)

func SetController(ctrl *controller.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ctrl", ctrl)
		c.Next()
	}
}

func GetController(c *gin.Context) *controller.Controller {
	return c.MustGet("ctrl").(*controller.Controller)
}
