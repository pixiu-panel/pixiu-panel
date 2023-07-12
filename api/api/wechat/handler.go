package wechat

import (
	"context"
	"encoding/xml"
	"fmt"
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/redis"
	"pixiu-panel/internal/wechat"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/constant"
	"time"
)

// handler
// @description: 处理微信消息回调
// @param msg
func handler(msg param.WechatCallback) {
	switch msg.Type {
	case constant.WechatMessageTypeAddFriend:
		// 添加好友
		addFriendHandler(msg)
	}
}

// addFriendHandler
// @description: 添加好友
// @param msg
func addFriendHandler(msg param.WechatCallback) {
	// 解析xml消息
	var addFriend param.WechatAddFriend
	if err := xml.Unmarshal([]byte(msg.Content), &addFriend); err != nil {
		log.Errorf("解析添加好友消息失败: %v", err)
		return
	}
	// 组装redisKey
	rdsKey := fmt.Sprintf("notify:bind:waiting:%s", addFriend.Content)
	// 如果不存在，直接返回
	if has, _ := redis.Client.Exists(context.Background(), rdsKey).Result(); has == 0 {
		return
	}
	// 存在，同意请求
	if err := wechat.AcceptAddFriend(addFriend.EncryptUsername, addFriend.Ticket); err != nil {
		// 同意好友请求失败
		log.Errorf("同意好友请求失败: %v", err)
		return
	}
	// 成功，修改Redis缓存数据，设置五分钟内过期
	cacheMsg := fmt.Sprintf("{\"account\": \"%s\",\"nickname\":\"%s\",\"type\":\"wechat\"}", addFriend.FromUsername, addFriend.FromNickname)
	if err := redis.Client.Set(context.Background(), rdsKey, cacheMsg, 5*time.Minute).Err(); err != nil {
		log.Errorf("修改Redis缓存数据失败: %v", err)
	}
}
