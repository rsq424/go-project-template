package router

import (
	"github.com/gin-gonic/gin"

	"nenly-wechat-api/middleware"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	PublicGroup := Router.Group("")
	{
		InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	return Router
}
