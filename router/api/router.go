package api

import "github.com/gin-gonic/gin"

// InitRouter
// @description: 初始化路由
func InitRouter(g *gin.RouterGroup) {
	notify(g.Group("/notify")) // 消息通知
}
