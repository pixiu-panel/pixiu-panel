package initialize

import (
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
)

// databaseTable
// @description: 初始化数据库表
func databaseTable() {
	tables := []any{
		new(entity.User),           // 用户表
		new(entity.UserNotify),     // 用户推送配置表
		new(entity.UserJd),         // 用户京东账号绑定表
		new(entity.NotifyLog),      // 推送记录表
		new(entity.InvitationCode), // 邀请码
	}

	// 同步表结构
	if err := db.Client.AutoMigrate(tables...); err != nil {
		log.Panicf("初始化数据库表失败: %v", err)
	}
}
