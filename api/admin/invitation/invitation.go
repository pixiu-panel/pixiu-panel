package invitation

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/invitation"
)

// Page
// @description: 分页查询邀请码
// @param ctx
func Page(ctx *gin.Context) {
	var p param.PageInvitationCode
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}
	if p.UserId == "mine" {
		p.UserId = ctx.GetString("userId")
	}

	records, total, err := invitation.Page(p)
	if err != nil {
		response.New(ctx).SetMsg("查询失败").SetError(err).Fail()
		return
	}
	// 返回数据
	response.New(ctx).SetData(response.NewPageData(records, total, p.Current, p.Size)).Success()
}
