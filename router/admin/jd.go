package admin

import (
	"github.com/gin-gonic/gin"
	jdApi "pixiu-panel/api/admin/jd"
)

// jd
// @description: 京东相关接口
// @param g
func jd(g *gin.RouterGroup) {
	bind := g.Group("/binding")
	bind.GET("", jdApi.GetBind)                   // 获取用户绑定的京东账号
	bind.DELETE("", jdApi.Delete)                 // 删除绑定的京东账号
	bind.GET("/qrcode", jdApi.GetJdQrcode)        // 获取京东二维码
	bind.GET("/qrcode/check", jdApi.CheckBinding) // 获取京东二维码扫描状态
}
