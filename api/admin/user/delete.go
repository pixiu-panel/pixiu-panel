package user

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/user"
	"unicode/utf8"
)

// Delete
// @description: 删除用户
// @param ctx
func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if utf8.RuneCountInString(id) != 32 {
		response.New(ctx).SetMsg("参数错误").Fail()
		return
	}

	if err := user.Delete(id); err != nil {
		response.New(ctx).SetMsg("删除失败").SetError(err).Fail()
		return
	}
	response.New(ctx).Success()
}
