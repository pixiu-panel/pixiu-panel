package notify

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/notify"
)

// PageNotifyLog
// @description: 分页获取消息通知日志
// @param ctx
func PageNotifyLog(ctx *gin.Context) {
	var p param.PageNotifyLog
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}
	// 补充参数
	if p.UserId == "mine" {
		// 查询自己的
		p.UserId = ctx.GetString("userId")
	}

	// 查询数据
	records, total, err := notify.PageNotifyLog(p)
	if err != nil {
		log.Errorf("获取消息通知日志失败: %v", err)
		response.New(ctx).SetMsg("数据获取失败").SetError(err).Fail()
		return
	}
	// 返回数据
	response.New(ctx).SetData(records).SetData(response.NewPageData(records, total, p.Current, p.Size)).Success()
}
