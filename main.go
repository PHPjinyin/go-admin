package main

import (
	"jin-gin/bootstrap"
	"jin-gin/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志
	bootstrap.InitializeLog()
	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 初始化校验函数
	bootstrap.InitializeValidator()
	// 初始化 Redis
	bootstrap.InitializeRedis()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	global.App.Log.Info("log init success!")
	bootstrap.RunServer()
}
