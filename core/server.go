package core

import (
	"nenly-wechat-api/router"
)

func RunWindowsServer() {
	Router := router.Routers()
	Router.Run(":9002")
}