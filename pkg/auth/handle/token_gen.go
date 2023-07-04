package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/google/uuid"
	"pixiu-panel/internal/redis"
	"strings"
)

func NewAccessGenerate() *AccessGenerate {
	return &AccessGenerate{}
}

type AccessGenerate struct {
}

// Token
// @description: 手动实现Token生成，直接生成UUID，替换掉自带的那个憨得一批的长长的字符串
// @receiver ag
// @param ctx context.Context 上下文
// @param data *oauth2.GenerateBasic 生成基础数据
// @param isGenRefresh bool 是否生成RefreshToken
// @return string AccessToken 令牌
// @return string RefreshToken 刷新令牌
// @return error 错误信息
func (ag *AccessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	u, _ := uuid.NewUUID()
	access := strings.ReplaceAll(u.String(), "-", "")

	refresh := ""
	if isGenRefresh {
		u, _ = uuid.NewUUID()
		refresh = strings.ReplaceAll(u.String(), "-", "")
	}
	// 生成新的，清理掉旧的
	ag.clearOldToken(ctx, data.UserID)
	// 返回结果
	return access, refresh, nil
}

// clearOldToken
// @description: 清理掉旧的Token和RefreshToken
// @receiver ag
// @param ctx context.Context 上下文
// @param userId string 用户ID
func (ag *AccessGenerate) clearOldToken(ctx context.Context, userId string) {
	key := fmt.Sprintf("oauth:token:token:%s", userId)
	accessToken, err := redis.Client.Get(context.Background(), key).Result()
	if err != nil {
		//log.Errorf("获取缓存用户的accessToken失败: %v", err.Error())
		return
	}
	if accessToken != "" {
		// 老的Token
		var baseKey string
		baseKey, err = redis.Client.Get(ctx, "oauth:token:"+accessToken).Result()
		if err != nil {
			return
		}
		// 老Token详细数据
		var dataStr string
		dataStr, err = redis.Client.Get(ctx, "oauth:token:"+baseKey).Result()
		if err != nil {
			return
		}
		var m map[string]interface{}
		if err = json.Unmarshal([]byte(dataStr), &m); err != nil {
			return
		}
		// 删除AccessToken等信息
		redis.Client.Del(ctx, fmt.Sprintf("oauth:token:%v", m["Access"]))
		redis.Client.Del(ctx, fmt.Sprintf("oauth:token:%v", m["Refresh"]))
		redis.Client.Del(ctx, "oauth:token:"+baseKey)
	}
}
