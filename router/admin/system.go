package admin

import (
	"github.com/gin-gonic/gin"
	systemApi "pixiu-panel/api/admin/system"
)

// system
// @description: 系统相关接口
// @param g
func system(g *gin.RouterGroup) {
	g.GET("/notify", systemApi.GetNotifyConfig)
}
