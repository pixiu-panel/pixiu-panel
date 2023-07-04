package tasks

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/go-co-op/gocron"
	"time"
)

// StartScheduled
// @description: 启动定时任务
func StartScheduled() {
	// 定时任务发送消息
	s := gocron.NewScheduler(time.Local)
	_, _ = s.Every("5m").Do(updateJdAccount) // 每5分钟检测一次京东账号状态
	// 开启定时任务
	s.StartAsync()
	log.Debugf("定时任务初始化成功")
}
