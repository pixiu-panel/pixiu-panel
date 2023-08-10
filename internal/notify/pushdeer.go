package notify

import (
	"errors"
	"gitee.ltd/lxh/logger/log"
	"github.com/go-resty/resty/v2"
	"pixiu-panel/config"
	"strings"
)

// pushDeer
// @description: PushDeer推送
type pushDeer struct {
	key string // 消息推送key
}

// Send
// @description: 发送通知
// @receiver s
// @param title
// @param content
// @return err
func (s pushDeer) Send(title, content string) (err error) {
	// 判断是否启用
	if !config.Conf.Notify.PushDeer.Enable {
		err = errors.New("PushDeer推送未启用")
		return
	}
	// 组装参数
	param := map[string]any{
		"text":    title,
		"desp":    content,
		"type":    "markdown",
		"pushkey": s.key,
	}

	// 处理接口地址
	api := "https://api2.pushdeer.com/message/push"
	if config.Conf.Notify.PushDeer.Host != "" && strings.Index(config.Conf.Notify.PushDeer.Host, "http") != -1 {
		api = config.Conf.Notify.PushDeer.Host
	}

	// 发送消息
	client := resty.New()
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post(api)
	if err != nil {
		log.Errorf("[PushDeer]消息发送失败: %s", err.Error())
	}

	return
}
