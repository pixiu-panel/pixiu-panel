package notify

import (
	"encoding/json"
)

// send
// @description: 通知发送接口
type send interface {
	Send(string, string) error // 发送通知(title, content)
}

// New
// @description: 发送通知
// @param channel 通知渠道
// @param params 通知参数
// @return send 通知发送接口
func New(channel, param string) send {
	switch channel {
	case "wechat":
		// 微信通知
		return wechat{param}
	case "qq":
		// QQ通知
		return qq{param}
	case "ftqq":
		// Server酱通知
		return ftqq{param}
	case "pushdeer":
		// PushDeer通知
		return pushDeer{param}
	case "pushplus":
		// PushPlus通知
		// 解析配置为结构体
		var pp pushPlus
		_ = json.Unmarshal([]byte(param), &pp)
		return pp
	case "smtp":
		// 邮件通知
		return email{param}
	default:
		return unknown{param}
	}
}
