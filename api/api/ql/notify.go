package ql

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gin-gonic/gin"
	"log"
	"pixiu-panel/config"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/notify"
	"strings"
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

	// 处理一下部分消息内容
	if len(config.Conf.Notify.Replace) > 0 {
		for _, rule := range config.Conf.Notify.Replace {
			if p.Title == rule.Title {
				p.Content = strings.ReplaceAll(p.Content, rule.Source, rule.Destination)
			}
		}
	}

	go func() {
		// 异步保存原始日志入库
		var rl entity.RawNotifyLog
		rl.Title = p.Title
		rl.Content = p.Content
		if err := db.Client.Create(&rl).Error; err != nil {
			log.Printf("保存原始日志失败: %v", err)
		}
	}()

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
