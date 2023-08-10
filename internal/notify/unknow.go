package notify

import (
	"errors"
	"gitee.ltd/lxh/logger/log"
)

// unknown
// @description: 未知通知
type unknown struct {
	Param string // 通知参数
}

// Send
// @description: 发送通知
// @receiver s
// @param title
// @param content
// @return err
func (s unknown) Send(title, content string) (err error) {
	log.Debugf("未知通知: %s --> %s", title, content)
	return errors.New("未知的通知类型")
}
