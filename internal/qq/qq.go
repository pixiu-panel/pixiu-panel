package qq

import (
	"github.com/go-resty/resty/v2"
	"pixiu-panel/config"
)

// AcceptAddFriend
// @description: 同意QQ好友申请
// @param flag
// @return err
func AcceptAddFriend(flag string) (err error) {
	// 组装参数
	param := map[string]any{
		"flag":    flag, // 加好友请求的 flag（需从上报的数据中获得）
		"approve": true, // 是否同意请求
		"remark":  "",   // 添加后的好友备注（仅在同意时有效）
	}

	client := resty.New()
	_, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post(config.Conf.Notify.QQ.Host + "/set_friend_add_request")
	return
}
