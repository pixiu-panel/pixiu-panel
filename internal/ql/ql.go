package ql

import (
	"context"
	"github.com/go-resty/resty/v2"
	"pixiu-panel/config"
	"pixiu-panel/internal/ql/model"
	"pixiu-panel/internal/redis"
	"time"
)

// GetToken
// @description: 获取青龙Token
// @return token
// @return err
func GetToken() (token model.Token, err error) {
	// 先取缓存，不存在再调接口
	rdsKey := "ql:token"
	if redis.Client.Get(context.Background(), rdsKey).Scan(&token) == nil {
		return
	}
	// 从接口取
	if token, err = getToken(); err != nil {
		return
	}

	// 计算当前时间到过期时间的秒数再减去一天，作为缓存时间
	token.Expiration = token.Expiration - int(time.Now().Unix()) - 86400

	// 缓存一下数据
	expiration := time.Second * time.Duration(token.Expiration)
	err = redis.Client.Set(context.Background(), rdsKey, &token, expiration).Err()
	return
}

// getToken
// @description: 获取青龙Token
// @return token
// @return err
func getToken() (token model.Token, err error) {
	var resp model.Response[model.Token]
	cli := resty.New()
	_, err = cli.R().
		SetQueryParams(map[string]string{
			"client_id":     config.Conf.Ql.ClientId,
			"client_secret": config.Conf.Ql.ClientSecret,
		}).
		SetResult(&resp).
		Get(config.Conf.Ql.Host + "/open/auth/token")
	if err == nil {
		token = resp.Data
	}
	return
}
