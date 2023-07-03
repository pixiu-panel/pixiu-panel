package login

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"pixiu-panel/config"
	"pixiu-panel/pkg/response"
	"time"
)

// Login
// @description: 登录
// @param ctx
// @return err
func Login(ctx echo.Context) (err error) {
	log.Debugf("收到登录请求")
	var p LoginWithPassword
	if err = ctx.Bind(&p); err != nil {
		log.Errorf("参数解析失败：%v", err)
		return response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
	}

	// 校验密码
	if p.Username != "lxh" || p.Password != "admin123" {
		return response.New(ctx).SetMsg("用户名或密码错误").Fail()
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"张三丰",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	// 生成Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成token
	var tokenStr string
	tokenStr, err = token.SignedString([]byte(config.Conf.System.Jwt.Secret))
	if err != nil {
		return err
	}
	return response.New(ctx).SetData(tokenStr).Success()
}
