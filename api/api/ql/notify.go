package ql

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"log"
	"pixiu-panel/config"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/notify"
)

// MessageNotify
// @description: 青龙消息通知
// @param ctx
func MessageNotify(ctx *gin.Context) {
	token := ctx.Query("token")
	log.Printf("来源: %s", token)

	// 获取body数据
	var p param.NotifyMessage
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}
	log.Printf("当前处理任务: %s", p.Title)

	// 校验任务是否需要发送消息
	if !slice.Contain(config.Conf.Notify.AllowTitle, p.Title) {
		log.Printf("当前任务不需要发送消息，原始消息内容:\n%s", p.Content)
		response.New(ctx).Success()
		return
	}

	// 处理消息
	if err := notify.Parse(p); err != nil {
		response.New(ctx).SetMsg("消息处理失败").SetError(err).Fail()
		return
	}

	// 返回成功
	response.New(ctx).SetMsg("已抄收").Success()
}
