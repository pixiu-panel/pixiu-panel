package admin

import (
	"github.com/labstack/echo/v4"
	"pixiu-panel/router/middleware"
)

// InitRouter
// @description: 初始化路由
func InitRouter(g *echo.Group) {
	login(g.Group("/login")) // 登录相关接口
	g.Use(middleware.Jwt())  // 下面的接口需要登录才能访问
	user(g.Group("/user"))   // 用户相关接口
}
