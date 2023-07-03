package user

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"pixiu-panel/model/cache"
	"pixiu-panel/pkg/response"
)

// Info
// @description: 用户信息
// @param ctx
// @return err
func Info(ctx echo.Context) (err error) {
	log.Debugf("收到获取用户信息请求")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*cache.JwtCustomClaims)
	name := claims.Username

	return response.New(ctx).SetData(name).Success()
}
