package notify

import (
	"errors"
	"gitee.ltd/lxh/logger/log"
	"github.com/go-resty/resty/v2"
	"pixiu-panel/config"
	"strings"
)

// ftqq
// @description: 方糖气球通知
type ftqq struct {
	key string // key
}

// Send
// @description: 发送通知
// @receiver s
// @param title
// @param content
// @return err
func (s ftqq) Send(title, content string) (err error) {
	// 判断是否启用
	if !config.Conf.Notify.Ftqq.Enable {
		err = errors.New("server酱通知未启用")
		return
	}
	// 判断Key是否包含SCT，如果是，使用新版接口
	apiHost := ""
	if strings.Index(s.key, "SCT") != -1 {
		apiHost = "https://sctapi.ftqq.com/"
	} else {
		apiHost = "https://sc.ftqq.com/"
	}

	// 组装参数
	param := map[string]any{
		"text": title,
		"desp": strings.ReplaceAll(content, "\n", "\n\n"),
	}

	// 发送消息
	client := resty.New()
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post(apiHost + s.key + ".send")
	if err != nil {
		log.Errorf("[Server酱]消息发送失败: %s", err.Error())
	}
	return
}
