package main

import (
	"nenly-wechat-api/config"
	"nenly-wechat-api/core"
	"nenly-wechat-api/global"

	"github.com/gin-gonic/gin"

	"time"
)

func main() {
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc

	err := config.Setup()
	if err != nil{
		panic(any(err))
	}

	global.Config = config.GetConfig()

	// 设置debug模式
	if config.GetConfig().System.SystemEnv != "test" {
		gin.SetMode(gin.ReleaseMode)
	}

	core.RunWindowsServer()
}
