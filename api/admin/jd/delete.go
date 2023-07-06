package jd

import (
	"github.com/kataras/iris/v12"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/jd"
)

// Delete
// @description: 删除京东账号
// @param ctx
func Delete(ctx iris.Context) {
	// 获取参数
	userId := ctx.Value("userId").(string)
	id := ctx.URLParam("id")

	// 删除数据
	if err := jd.Delete(userId, id); err != nil {
		response.New(ctx).SetMsg("删除失败").SetError(err).Fail()
		return
	}
	// 返回数据
	response.New(ctx).Success()
}
