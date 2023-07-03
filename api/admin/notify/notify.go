package notify

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"pixiu-panel/model/cache"
	"pixiu-panel/pkg/response"
)

// Binding
// @description: 绑定推送渠道
// @param ctx
// @return err
func Binding(ctx iris.Context) {
	var p BindingNotify
	if err := ctx.ReadBody(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}

	// 取出登录用户Id
	claims := jwt.Get(ctx).(*cache.JwtCustomClaims)
	userId := claims.Id
	log.Debugf("收到绑定推送渠道请求，用户Id：%d", userId)

	// 根据传入类型返回对应数据

	return
}
