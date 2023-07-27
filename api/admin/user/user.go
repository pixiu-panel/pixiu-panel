package user

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/user"
)

// Page
// @description: 用户列表
// @param ctx
// @return err
func Page(ctx *gin.Context) {
	var p param.PageUser
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}

	records, total, err := user.Page(p)
	if err != nil {
		response.New(ctx).SetMsg("数据获取失败").SetError(err).Fail()
		return
	}

	response.New(ctx).SetData(response.NewPageData(records, total, p.Current, p.Size)).Success()
}
