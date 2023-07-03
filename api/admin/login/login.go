package login

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"pixiu-panel/config"
	"pixiu-panel/model/cache"
	"pixiu-panel/pkg/response"
	userService "pixiu-panel/service/user"
	"pixiu-panel/utils"
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
	userInfo, err := userService.GetUserWithLogin(p.Username)
	if err != nil {
		return response.New(ctx).SetMsg("账号不存在").Fail()
	}
	// 校验密码
	if !utils.PasswordUtils().ComparePassword(userInfo.Password, p.Password) {
		return response.New(ctx).SetMsg("密码错误").Fail()
	}
	defer func() {
		// 如果登录成功，更新登录时间
		if err == nil {
			userService.UpdateLastLoginInfo(userInfo.Id, ctx.RealIP())
		}
	}()

	// 设置JWT携带的信息
	claims := &cache.JwtCustomClaims{
		Id:         userInfo.Id,
		Username:   userInfo.Username,
		IsVerified: userInfo.IsVerified,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	// 生成Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成token字符串
	var tokenStr string
	tokenStr, err = token.SignedString([]byte(config.Conf.System.Jwt.Secret))
	if err != nil {
		return err
	}
	return response.New(ctx).SetData(tokenStr).Success()
}
