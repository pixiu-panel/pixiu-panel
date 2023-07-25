package admin

import (
	"github.com/gin-gonic/gin"
	invitationApi "pixiu-panel/api/admin/invitation"
)

// invitation
// @description: 邀请码管理
// @param g
func invitation(g *gin.RouterGroup) {
	g.GET("", invitationApi.Page)     // 分页查询邀请码
	g.POST("/new", invitationApi.Gen) // 新增邀请码
}
