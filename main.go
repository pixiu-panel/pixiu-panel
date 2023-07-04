package main

import (
	"gitee.ltd/lxh/logger"
	"gitee.ltd/lxh/logger/log"
	"github.com/kataras/iris/v12"
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
	e := iris.Default()
	validator.Init(e)

	// 初始化后台路由
	admin.InitRouter(e.Party("/admin/v1"))
	// 初始化开放接口路由
	api.InitRouter(e.Party("/api/v1"))

	// 启动服务
	if err := e.Listen(":1323"); err != nil {
		log.Errorf("服务启动失败：%v", err)
		return
	}
}
