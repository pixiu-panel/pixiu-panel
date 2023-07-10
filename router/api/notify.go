package api

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/api/api/ql"
	"pixiu-panel/api/api/wechat"
)

// notify
// @description: 青龙消息通知
// @param g
func notify(g *gin.RouterGroup) {
	g.POST("/message", ql.MessageNotify) // 消息通知接口(模拟实现 gotify)
	g.POST("/wechat", wechat.Notify)     // 微信消息回调
}
