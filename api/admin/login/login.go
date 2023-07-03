package login

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
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
func Login(ctx iris.Context) {
	log.Debugf("收到登录请求")
	var p LoginWithPassword
	if err := ctx.ReadJSON(&p); err != nil {
		log.Errorf("参数解析失败：%v", err)
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}

	// 校验密码
	userInfo, err := userService.GetUserWithLogin(p.Username)
	if err != nil {
		response.New(ctx).SetMsg("账号不存在").Fail()
		return
	}
	// 校验密码
	if !utils.PasswordUtils().ComparePassword(userInfo.Password, p.Password) {
		response.New(ctx).SetMsg("密码错误").Fail()
		return
	}
	defer func() {
		// 如果登录成功，更新登录时间
		if err == nil {
			userService.UpdateLastLoginInfo(userInfo.Id, ctx.RemoteAddr())
		}
	}()

	// 设置JWT携带的信息
	claims := &cache.JwtCustomClaims{
		Id:         userInfo.Id,
		Username:   userInfo.Username,
		IsVerified: userInfo.IsVerified,
	}

	signer := jwt.NewSigner(jwt.HS256, config.Conf.System.Jwt.Secret, 10*time.Minute)

	// 生成Token对象
	var token []byte
	token, err = signer.Sign(claims)
	if err != nil {
		response.New(ctx).SetMsg("系统错误").Fail()
		return
	}
	response.New(ctx).SetData(string(token)).Success()
}
