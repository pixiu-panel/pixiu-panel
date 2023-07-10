package main

import (
	"gitee.ltd/lxh/logger"
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"pixiu-panel/internal/initialize"
	"pixiu-panel/internal/wechat"
	"pixiu-panel/pkg/validator"
	"pixiu-panel/router/admin"
	"pixiu-panel/router/api"
	"syscall"
)

// init
// @description: 初始化系统
func init() {
	// 初始化日志工具
	logger.InitLogger(logger.LogConfig{Mode: logger.Dev, LokiEnable: false, FileEnable: true})
	// 初始化系统
	initialize.InitSystem()
	// 配置微信消息回调
	wechat.SetCallback()
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

	go func() {
		// 启动服务
		if err := app.Run(":1323"); err != nil {
			log.Errorf("服务启动失败：%v", err)
			return
		}
	}()

	// 监控两个信号
	// TERM信号（kill + 进程号 触发）
	// 中断信号（ctrl + c 触发）
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	s := <-osc
	log.Debugf("监听到退出信号,s=%v，开始执行退出前操作", s)
	wechat.ClearCallback()
}
