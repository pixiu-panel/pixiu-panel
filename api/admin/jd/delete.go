package jd

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/jd"
)

// Delete
// @description: 删除京东账号
// @param ctx
func Delete(ctx *gin.Context) {
	// 获取参数
	userId := ctx.GetString("userId")
	id := ctx.Param("id")

	// 删除数据
	if err := jd.Delete(userId, id); err != nil {
		response.New(ctx).SetMsg("删除失败").SetError(err).Fail()
		return
	}
	// 返回数据
	response.New(ctx).Success()
}
