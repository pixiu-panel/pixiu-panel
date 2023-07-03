package main

import (
	"gitee.ltd/lxh/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pixiu-panel/internal/initialize"
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
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 初始化后台路由
	admin.InitRouter(e.Group("/admin/v1"))
	// 初始化开放接口路由
	api.InitRouter(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(":1323"))
}
