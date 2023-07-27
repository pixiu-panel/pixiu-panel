package admin

import (
	"github.com/gin-gonic/gin"
	userApi "pixiu-panel/api/admin/user"
)

// user
// @description: 用户相关接口
// @param g
func user(g *gin.RouterGroup) {
	g.POST("/password", userApi.ChangePassword) // 修改密码
	g.GET("/page", userApi.Page)                // 获取用户列表
	g.DELETE("/:id", userApi.Delete)            // 删除用户
}
