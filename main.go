package main

import (
	"gitee.ltd/lxh/logger"
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/internal/initialize"
	"pixiu-panel/pkg/validator"
	"pixiu-panel/router/admin"
	"pixiu-panel/router/api"
)

// init
// @description: 初始化系统
func init() {
	// 初始化日志工具
	logger.InitLogger(logger.LogConfig{Mode: logger.Dev, LokiEnable: false, FileEnable: true})
	// 初始化系统
	initialize.InitSystem()
}

// main
// @description: 启动入口
func main() {
	// 注册参数绑定错误信息翻译器
	validator.Init()
	app := gin.Default()

	// 初始化后台路由
	admin.InitRouter(app.Group("/admin/v1"))
	// 初始化开放接口路由
	api.InitRouter(app.Group("/api/v1"))

	// 启动服务
	if err := app.Run(":1323"); err != nil {
		log.Errorf("服务启动失败：%v", err)
		return
	}
}
