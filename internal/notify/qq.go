package notify

import (
	"errors"
	"fmt"
	"pixiu-panel/config"
	qqService "pixiu-panel/internal/qq"
)

// qq
// @description: QQ通知
type qq struct {
	toUser string // 接收人
}

// Send
// @description: 发送通知
// @receiver s
// @param title
// @param content
// @return err
func (s qq) Send(title, content string) (err error) {
	// 判断是否启用
	if !config.Conf.Notify.QQ.Enable {
		err = errors.New("QQ通知未启用")
		return
	}
	// 组装消息
	msg := fmt.Sprintf("%s\n \n%s", title, content)
	// 发送消息
	return qqService.SendMessage(s.toUser, msg)
}
