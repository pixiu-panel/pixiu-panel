package auth

import (
	"context"
	"encoding/json"
	"gitee.ltd/lxh/logger/log"
	"github.com/go-oauth2/oauth2/v4/models"
	"pixiu-panel/internal/redis"
)

// GetUserIdWithRefreshToken
// @description: 根据refreshToken获取userId
// @param refreshToken string refreshToken
// @return userId string userId
func GetUserIdWithRefreshToken(refreshToken string) (userId string) {
	// 取出真实保存token信息的key
	realKey, err := redis.Client.Get(context.Background(), "oauth:token:"+refreshToken).Result()
	if err != nil {
		log.Errorf("根据refreshToken获取realKey失败: %v", err)
		return
	}
	// 取出缓存信息
	var (
		tiStr     string       // 保存的原始字符串
		tokenInfo models.Token // Token结构体
	)
	tiStr, err = redis.Client.Get(context.Background(), "oauth:token:"+realKey).Result()
	if err != nil {
		log.Errorf("获取tokenInfo失败: %v", err)
		return
	}
	// 反序列化
	if err = json.Unmarshal([]byte(tiStr), &tokenInfo); err != nil {
		log.Errorf("tokenInfo反序列化失败: %v", err)
		return
	}
	// 返回userId
	return tokenInfo.GetUserID()
}
