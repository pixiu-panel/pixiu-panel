package admin

import (
	"github.com/gin-gonic/gin"
	menuApi "pixiu-panel/api/admin/menu"
)

// menu
// @description: 菜单相关接口
// @param g
func menu(g *gin.RouterGroup) {
	g.GET("", menuApi.GetMenuTree)
}
