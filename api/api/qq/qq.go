package qq

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/labstack/echo/v4"
)

// Notify
// @description: QQ消息回调
// @param ctx
// @return err
func Notify(ctx echo.Context) (err error) {
	log.Debugf("收到QQ消息回调")
	return
}
