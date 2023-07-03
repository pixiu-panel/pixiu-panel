package admin

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"pixiu-panel/config"
	"pixiu-panel/pkg/response"
)

// InitRouter
// @description: 初始化路由
func InitRouter(g *echo.Group) {
	login(g.Group("/login")) // 登录相关接口
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: config.Conf.System.Jwt.Secret,
		ErrorHandler: func(ctx echo.Context, err error) error {
			log.Debugf("未授权访问: %v", err)
			return response.New(ctx).SetCode(http.StatusUnauthorized).SetMsg("未登录或登录已过期").Fail()
		},
	})) // 下面的接口需要登录才能访问
	user(g.Group("/user")) // 用户相关接口
}
