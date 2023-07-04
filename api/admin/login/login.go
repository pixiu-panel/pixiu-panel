package login

import (
	"context"
	"encoding/json"
	"gitee.ltd/lxh/logger/log"
	"github.com/kataras/iris/v12"
	"net/http"
	"net/url"
	"pixiu-panel/internal/redis"
	"pixiu-panel/pkg/auth"
	"pixiu-panel/pkg/response"
	"pixiu-panel/pkg/validator"
	userService "pixiu-panel/service/user"
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
		response.New(ctx).SetMsg("参数错误").SetError(validator.Translate(err)).Fail()
		return
	}

	// 重写参数
	ctx.Request().Form = url.Values{
		"username":   {p.Username},
		"password":   {p.Password},
		"scope":      {"ALL"},
		"grant_type": {"password"},
	}
	// 参数解析成功，进行登录
	if err := auth.OAuthServer.HandleTokenRequest(ctx.ResponseWriter(), ctx.Request()); err != nil {
		log.Errorf("登录失败：%s", err.Error())
		response.New(ctx).SetMsg("系统错误").SetError(err).Fail()
		return
	}

	if ctx.ResponseWriter().StatusCode() == http.StatusOK {
		go userService.UpdateLastLoginInfo(p.Username, ctx.RemoteAddr())
	}
}

// Refresh
// @description: 刷新Token
// @param ctx
func Refresh(ctx iris.Context) {
	var p RefreshToken
	if err := ctx.ReadBody(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(validator.Translate(err)).Fail()
		return
	}

	// 取出用户Id
	userId := auth.GetUserIdWithRefreshToken(p.RefreshToken)

	// 重写参数
	ctx.Request().Form = url.Values{
		"refresh_token": {p.RefreshToken},
		"grant_type":    {"refresh_token"},
	}

	// 刷新Token
	if err := auth.OAuthServer.HandleTokenRequest(ctx.ResponseWriter(), ctx.Request()); err != nil {
		log.Errorf("Token数据返回失败: %v", err.Error())
		response.New(ctx).SetMsg("系统错误").Fail()
	}

	// 登录成功才更新登录时间
	if ctx.ResponseWriter().StatusCode() == http.StatusOK {
		// 登录成功，更新登录时间和IP
		go userService.UpdateLastLoginInfo(userId, ctx.RemoteAddr())
	}
}

// Logout
// @description: 退出登录
// @param ctx
func Logout(ctx iris.Context) {
	log.Debug("退出登录啦")
	// Token字符串前缀
	const bearerSchema string = "Bearer "
	// 取出Token
	tokenHeader := ctx.GetHeader("Authorization")
	tokenStr := tokenHeader[len(bearerSchema):]
	// 取出原始RedisKey
	baseDataId, err := redis.Client.Get(context.Background(), "oauth:token:"+tokenStr).Result()
	if err != nil {
		response.New(ctx).SetMsg("Token信息获取失败").Fail()
		return
	}
	baseDataStr, err := redis.Client.Get(context.Background(), "oauth:token:"+baseDataId).Result()
	if err != nil {
		response.New(ctx).SetMsg("Token信息获取失败").Fail()
		return
	}
	// 转换数据为Map
	tokenData := make(map[string]interface{})
	if err = json.Unmarshal([]byte(baseDataStr), &tokenData); err != nil {
		response.New(ctx).SetMsg("系统错误").SetError(err).Fail()
		return
	}
	// 删除Redis缓存的数据
	redis.Client.Del(context.Background(), "oauth:token:"+baseDataId)
	redis.Client.Del(context.Background(), "oauth:token:"+tokenData["Access"].(string))
	redis.Client.Del(context.Background(), "oauth:token:"+tokenData["Refresh"].(string))

	response.New(ctx).Success()
}
