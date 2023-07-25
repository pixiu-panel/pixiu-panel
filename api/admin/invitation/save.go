package invitation

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
	"pixiu-panel/pkg/response"
	"pixiu-panel/utils"
)

// Gen
// @description: 生成邀请码
// @param ctx
func Gen(ctx *gin.Context) {
	// 构建数据
	var ent entity.InvitationCode
	ent.UserId = ctx.GetString("userId")
	ent.Code = utils.RandomUtils().GetRandomStringMini(8)
	// 入库
	if err := db.Client.Create(&ent).Error; err != nil {
		response.New(ctx).SetMsg("生成失败").SetError(err).Fail()
		return
	}
	response.New(ctx).Success()
}
