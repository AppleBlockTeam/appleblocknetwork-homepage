package main

import (
	"abnet_backend/source"
	"abnet_backend/source/logger"
	"abnet_backend/source/server"
)

func main() {
	// 初始化日志
	logger.InitLogger(logger.DEBUG)
	logger.Info("Hello!,Nya-ABNet-HomePage")

	// 初始化配置
	source.LoadConfig()

	// 启动Web服务器
	server.Setupserver()
}
