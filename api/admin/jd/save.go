package jd

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/jd"
)

// Update
// @description: 修改京东账号信息
// @param ctx
func Update(ctx *gin.Context) {
	var p param.SaveJdAccount
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}
	// 修改
	if err := jd.SaveJdInfo(p); err != nil {
		response.New(ctx).SetMsg("修改失败").SetError(err).Fail()
		return
	}
	response.New(ctx).SetMsg("修改成功").Success()
}
