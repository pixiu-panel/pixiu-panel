package login

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/config"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/user"
)

// Register
// @description: 注册
// @param ctx
func Register(ctx *gin.Context) {
	if !config.Conf.System.Register.Enable {
		response.New(ctx).SetMsg("未开启自主注册").Fail()
		return
	}
	// 解析参数
	var p param.Register
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}

	// 注册
	if err := user.Register(p); err != nil {
		response.New(ctx).SetMsg("注册失败").SetError(err).Fail()
		return
	}

	response.New(ctx).Success()
}
