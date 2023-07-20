package user

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/user"
)

// ChangePassword
// @description: 修改密码
// @param ctx
func ChangePassword(ctx *gin.Context) {
	var p param.ChangePassword
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}

	// 取出用户Id
	if p.UserId == "" {
		p.UserId = ctx.GetString("userId")
	}

	// 修改密码
	if err := user.ChangePassword(p); err != nil {
		response.New(ctx).SetMsg("修改失败").SetError(err).Fail()
		return
	}

	response.New(ctx).Success()
}
