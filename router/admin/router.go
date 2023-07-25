package admin

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/router/middleware"
)

// InitRouter
// @description: 初始化路由
func InitRouter(g *gin.RouterGroup) {
	login(g.Group("/login"))           // 登录相关接口
	g.Use(middleware.AuthorizeToken()) // 下面的接口需要登录才能访问
	menu(g.Group("/menu"))             // 菜单相关接口
	user(g.Group("/user"))             // 用户相关接口
	jd(g.Group("/jd"))                 // 京东相关接口
	notify(g.Group("/notify"))         // 通知相关接口
	invitation(g.Group("/invitation")) // 邀请码相关接口
}
