package main

import (
	"github.com/gin-gonic/gin"
	"jin-gin/bootstrap"
	"jin-gin/global"
	"net/http"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志
	bootstrap.InitializeLog()

	global.App.Log.Info("log init success!")

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// 启动服务器
	r.Run(":" + global.App.Config.App.Port)
}
