package login

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/config"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
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
	// 如果启用了邀请码，校验一下
	if config.Conf.System.Register.InvitationCode {
		log.Debugf("邀请码: %s", p.InvitationCode)
	}

	// 查询用户是否存在

	response.New(ctx).Success()
}
