package wechat

import (
	"fmt"
	"gitee.ltd/lxh/logger/log"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/go-resty/resty/v2"
	"net"
	"pixiu-panel/config"
)

// SetCallback
// @description: 设置微信消息回调
func SetCallback() {
	// 如果没配置微信推送渠道，则不执行
	if config.Conf.Notify.Wechat.Host == "" {
		return
	}
	// 获取本机ip
	ip := net.ParseIP(netutil.GetInternalIp())
	callback := fmt.Sprintf("http://%s:1323/api/v1/notify/wechat", ip.String())
	if config.Conf.Notify.Wechat.Callback != "" {
		callback = config.Conf.Notify.Wechat.Callback
	}
	log.Debugf("设置微信消息回调: %s", callback)
	// 组装参数
	param := map[string]any{
		"enableHttp": 1,           // 1表示使用http回调
		"ip":         ip.String(), // 本机ip
		"port":       "1323",      // Socket 端口，没啥用，因为用的是http回调，所以随便写个
		"url":        callback,    // 回调地址
		"timeOut":    5000,        // 超时时间(单位: 毫秒)
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("type", "9").
		SetBody(param).
		Post(config.Conf.Notify.Wechat.Host + "/api/")
	if err != nil {
		log.Panicf("设置微信消息回调失败: %s", err.Error())
	} else {
		log.Debugf("设置微信消息回调结果: %s", resp.String())
	}
}

// ClearCallback
// @description: 清除微信消息回调
func ClearCallback() {
	// 如果没配置微信推送渠道，则不执行
	if config.Conf.Notify.Wechat.Host == "" {
		return
	}

	// 调用接口
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("type", "10").
		Post(config.Conf.Notify.Wechat.Host + "/api/")

	if err != nil {
		log.Errorf("清除微信消息回调失败: %s", err.Error())
	} else {
		log.Debugf("清除微信消息回调结果: %s", resp.String())
	}
}

// AcceptAddFriend
// @description: 接受好友请求
// @param v3 string 结构体中的encryptusername字段
// @param v4 string 结构体中的ticket字段
// @return err error 错误
func AcceptAddFriend(v3, v4 string) (err error) {
	// 组装参数
	param := map[string]any{
		"v3":         v3,
		"vd":         v4,
		"permission": 8, // 仅聊天，0表示全部
	}

	// 调用接口
	client := resty.New()
	var resp *resty.Response
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		SetQueryParam("type", "23").
		Post(config.Conf.Notify.Wechat.Host + "/api/")

	if err != nil {
		log.Errorf("同意好友请求失败: %s", err.Error())
	} else {
		log.Debugf("同意好友请求成功: %s", resp.String())
	}

	return
}
