package notify

import (
	"errors"
	"fmt"
	"pixiu-panel/config"
	wechatService "pixiu-panel/internal/wechat"
)

// wechat
// @description: 微信通知
type wechat struct {
	wxId string // 接收人微信Id
}

// Send
// @description: 发送通知
// @receiver s
// @param title
// @param content
// @return err
func (s wechat) Send(title, content string) (err error) {
	// 判断是否启用
	if !config.Conf.Notify.Wechat.Enable {
		err = errors.New("微信通知未启用")
		return
	}
	// 组装消息
	msg := fmt.Sprintf("%s\n \n%s", title, content)
	// 发送消息
	return wechatService.SendMessage(s.wxId, msg)
}
