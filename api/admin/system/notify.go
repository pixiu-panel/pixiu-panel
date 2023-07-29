package system

import (
	"github.com/gin-gonic/gin"
	"pixiu-panel/config"
	"pixiu-panel/pkg/response"
)

// GetNotifyConfig
// @description: 获取推送渠道配置
// @param ctx
func GetNotifyConfig(ctx *gin.Context) {
	// 手动组装一下返回数据
	var data = make(map[string]map[string]any)
	data["wechat"] = map[string]any{
		"enable": config.Conf.Notify.Wechat.Enable,
		"qrcode": config.Conf.Notify.Wechat.QrCode,
	}
	data["qq"] = map[string]any{
		"enable": config.Conf.Notify.QQ.Enable,
		"qrcode": config.Conf.Notify.QQ.QrCode,
	}

	response.New(ctx).SetData(data).Success()
}
