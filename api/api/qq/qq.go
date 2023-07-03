package qq

import (
	"gitee.ltd/lxh/logger/log"
)

// Notify
// @description: QQ消息回调
// @param ctx
// @return err
func Notify(ctx iris.Context) (err error) {
	log.Debugf("收到QQ消息回调")
	return
}
