package qq

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
)

// Notify
// @description: QQ消息回调
// @param ctx
// @return err
func Notify(ctx *gin.Context) (err error) {
	log.Debugf("收到QQ消息回调")
	return
}
