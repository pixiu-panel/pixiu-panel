package notify

import (
	"errors"
	"gitee.ltd/lxh/logger/log"
	"github.com/go-resty/resty/v2"
	"pixiu-panel/config"
)

type pushPlus struct {
	token string // 接口请求Token
	user  string // 接收人
}

// Send
// @description: 发送通知
// @receiver s
// @param title
// @param content
// @return err
func (s pushPlus) Send(title, content string) (err error) {
	// 判断是否启用
	if !config.Conf.Notify.PushPlus.Enable {
		err = errors.New("PushPlus推送未启用")
		return
	}
	// 组装参数
	param := map[string]any{
		"token":   s.token,
		"title":   title,
		"content": content,
		"topic":   s.user,
	}

	// 发送消息
	client := resty.New()
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post("http://www.pushplus.plus/send")
	if err != nil {
		log.Errorf("[PushPlus]消息发送失败: %s", err.Error())
	}
	return
}
