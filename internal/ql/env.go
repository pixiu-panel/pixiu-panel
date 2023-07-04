package ql

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/go-resty/resty/v2"
	"pixiu-panel/config"
	"pixiu-panel/internal/ql/model"
)

// GetEnvs
// @description: 获取青龙环境变量
// @param key string 环境变量名
// @return envs []model.Env 环境变量列表
func GetEnvs(key string) (envs []model.Env) {
	token, err := GetToken()
	if err != nil {
		log.Errorf("获取青龙Token失败: %v", err)
		return
	}
	var resp model.Response[[]model.Env]
	cli := resty.New()
	_, err = cli.R().
		SetResult(&resp).
		SetQueryParam("searchValue", key).
		SetAuthToken(token.Token).
		Get(config.Conf.Ql.Host + "/open/envs")
	if err != nil {
		log.Errorf("获取青龙环境变量失败: %v", err)
		return
	}
	envs = resp.Data
	return
}
