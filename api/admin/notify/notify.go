package notify

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
)

// Binding
// @description: 绑定推送渠道
// @param ctx
// @return err
func Binding(ctx *gin.Context) {
	var p param.BindingNotify
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}

	// 取出登录用户Id
	userId := ctx.Value("userId").(string)
	log.Debugf("收到绑定推送渠道请求，用户Id：%d", userId)

	// 根据传入类型返回对应数据

	return
}
