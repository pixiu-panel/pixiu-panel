package tasks

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/go-co-op/gocron"
	"pixiu-panel/service/user"
	"time"
)

// StartScheduled
// @description: 启动定时任务
func StartScheduled() {
	// 定时任务发送消息
	s := gocron.NewScheduler(time.Local)

	// 每5分钟检测一次京东账号状态
	_, _ = s.Every("5m").Do(updateJdAccount)
	// 每天0点清理一次空的账号
	_, _ = s.Every(1).Days().At("00:00").Do(user.ClearEmptyUser)
	// 每天0点更新一次京东账号基础信息
	_, _ = s.Every(1).Days().At("00:00").Do(updateJdAccountInfo)

	// 开启定时任务
	s.StartAsync()
	log.Debugf("定时任务初始化成功")
}
