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

// DisableEnv
// @description: 禁用环境变量
// @param ids
// @return err
func DisableEnv(ids []int) (err error) {
	token, err := GetToken()
	if err != nil {
		log.Errorf("获取青龙Token失败: %v", err)
		return
	}

	cli := resty.New()
	_, err = cli.R().
		SetBody(ids).
		SetAuthToken(token.Token).
		Put(config.Conf.Ql.Host + "/open/envs/disable")
	if err != nil {
		log.Errorf("禁用青龙环境变量失败: %v", err)
	}
	return
}
