package qq

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/go-resty/resty/v2"
	"pixiu-panel/config"
)

// SendMessage
// @description: 发送QQ消息
// @param qq string QQ号
// @param msg string 消息内容
// @return err error 错误信息
func SendMessage(qq, msg string) (err error) {
	// 组装参数
	param := map[string]any{
		"user_id": qq,
		"message": msg,
	}

	client := resty.New()
	var resp *resty.Response
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post(config.Conf.Notify.QQ.Host + "/send_private_msg")
	log.Debugf("QQ消息发送结果: %s", resp.String())
	if err != nil {
		log.Errorf("给[%s]的消息发送失败: %s", qq, err.Error())
	}
	return
}
