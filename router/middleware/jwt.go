package middleware

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"pixiu-panel/config"
	"pixiu-panel/model/cache"
	"pixiu-panel/pkg/response"
)

// Jwt
// @description: Jwt中间件
func Jwt() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.Conf.System.Jwt.Secret),
		NewClaimsFunc: func(ctx echo.Context) jwt.Claims {
			return new(cache.JwtCustomClaims)
		},
		ErrorHandler: errorHandler,
	})
}

// errorHandler
// @description: 错误处理
// @param ctx
// @param err
// @return error
func errorHandler(ctx echo.Context, err error) error {
	log.Debugf("未授权访问: %v", err)
	return response.New(ctx).SetCode(http.StatusUnauthorized).SetMsg("未登录或登录已过期").Fail()
}
