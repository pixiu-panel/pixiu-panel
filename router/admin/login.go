package admin

import (
	"github.com/gin-gonic/gin"
	loginApi "pixiu-panel/api/admin/login"
)

// login
// @description: 登录相关接口
// @param g
func login(g *gin.RouterGroup) {
	g.POST("", loginApi.Login)           // 登录
	g.POST("/refresh", loginApi.Refresh) // 刷新Token
	g.POST("/logout", loginApi.Logout)   // 退出登录
}
