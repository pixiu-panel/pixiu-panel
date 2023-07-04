package redis

import (
	"context"
	"gitee.ltd/lxh/logger/log"
	"github.com/go-redis/redis/v8"
	"pixiu-panel/config"
)

var Client *redis.Client

// Init
// @description: 初始化redis客户端
func Init() {
	conf := config.Conf.Redis
	// 初始化连接
	conn := redis.NewClient(&redis.Options{
		Addr:     conf.GetDSN(),
		Password: conf.Password,
		DB:       conf.Db,
	})
	if err := conn.Ping(context.Background()).Err(); err != nil {
		log.Panicf("Redis连接初始化失败: %v", err)
	} else {
		log.Debug("Redis连接初始化成功")
	}
	Client = conn
}
