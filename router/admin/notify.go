package admin

import (
	"github.com/gin-gonic/gin"
	notifyApi "pixiu-panel/api/admin/notify"
)

// notify
// @description: 通知相关接口
// @param g
func notify(g *gin.RouterGroup) {
	g.POST("", notifyApi.Binding)           // 绑定微信预请求
	g.GET("", notifyApi.GetBindingAccounts) // 获取已绑定的推送渠道
	g.GET("/check", notifyApi.CheckBinding) // 检查绑定结果
	g.GET("/log", notifyApi.PageNotifyLog)  // 分页获取消息通知日志
}
