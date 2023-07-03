package wechat

import (
	"gitee.ltd/lxh/logger/log"
)

// Notify
// @description: 微信消息回调
// @param ctx
// @return err
func Notify(ctx iris.Context) (err error) {
	log.Debugf("收到微信消息回调")
	return
}
