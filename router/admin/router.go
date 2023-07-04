package admin

import (
	"github.com/kataras/iris/v12"
	"pixiu-panel/router/middleware"
)

// InitRouter
// @description: 初始化路由
func InitRouter(g iris.Party) {
	login(g.Party("/login"))           // 登录相关接口
	g.Use(middleware.AuthorizeToken()) // 下面的接口需要登录才能访问
	menu(g.Party("/menu"))             // 菜单相关接口
	user(g.Party("/user"))             // 用户相关接口
	jd(g.Party("/jd"))                 // 京东相关接口
}
