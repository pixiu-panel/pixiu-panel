package admin

import (
	"github.com/kataras/iris/v12"
	"pixiu-panel/router/middleware"
)

// InitRouter
// @description: 初始化路由
func InitRouter(g iris.Party) {
	login(g.Party("/login")) // 登录相关接口
	g.Use(middleware.Jwt())  // 下面的接口需要登录才能访问
	user(g.Party("/user"))   // 用户相关接口
}
