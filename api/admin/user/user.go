package user

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"pixiu-panel/model/cache"
	"pixiu-panel/pkg/response"
)

// Info
// @description: 用户信息
// @param ctx
// @return err
func Info(ctx iris.Context) {
	log.Debugf("收到获取用户信息请求")
	claims := jwt.Get(ctx).(*cache.JwtCustomClaims)
	name := claims.Username

	response.New(ctx).SetData(name).Success()
}
