package user

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/labstack/echo/v4"
	"pixiu-panel/pkg/response"
)

// Info
// @description: 用户信息
// @param ctx
// @return err
func Info(ctx echo.Context) (err error) {
	log.Debugf("收到获取用户信息请求")
	return response.New(ctx).Success()
}
