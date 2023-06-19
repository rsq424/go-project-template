package router

import (
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("")
	{
		BaseRouter.GET("/", func(c *gin.Context) {
			c.JSON(200,gin.H{"data":"hello world"})
		})
	}
	return BaseRouter
}

