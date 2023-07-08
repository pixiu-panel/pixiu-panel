package jd

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/jd"
)

// GetBind
// @description: 获取用户绑定的京东账号
// @param ctx
func GetBind(ctx *gin.Context) {
	var p param.PageUserJdAccount
	p.Current = -1
	p.Size = 10
	p.UserId = ctx.Value("userId").(string)

	// 查询数据
	records, _, err := jd.GetBindByUser(p)
	if err != nil {
		response.New(ctx).SetMsg("获取失败").SetError(err).Fail()
		return
	}
	// 返回数据
	response.New(ctx).SetData(records).Success()
}
