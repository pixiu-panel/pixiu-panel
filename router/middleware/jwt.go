package middleware

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	"net/http"
	"pixiu-panel/config"
	"pixiu-panel/model/cache"
	"pixiu-panel/pkg/response"
)

// Jwt
// @description: Jwt中间件
func Jwt() context.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, config.Conf.System.Jwt.Secret)
	verifier.WithDefaultBlocklist()

	// 自定义错误处理
	verifier.ErrorHandler = errorHandler

	// 返回验证器
	return verifier.Verify(func() interface{} {
		return new(cache.JwtCustomClaims)
	})
}

// errorHandler
// @description: 错误处理
// @param ctx
// @param err
// @return error
func errorHandler(ctx iris.Context, err error) {
	log.Debugf("未授权访问: %v", err)
	response.New(ctx).SetCode(http.StatusUnauthorized).SetMsg("未登录或登录已过期").Fail()
}
