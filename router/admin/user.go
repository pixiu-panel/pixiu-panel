package admin

import (
	"github.com/gin-gonic/gin"
	userApi "pixiu-panel/api/admin/user"
)

// user
// @description: 用户相关接口
// @param g
func user(g *gin.RouterGroup) {
	g.GET("", userApi.Info)                     // 登录用户信息
	g.POST("/password", userApi.ChangePassword) // 修改密码
}
