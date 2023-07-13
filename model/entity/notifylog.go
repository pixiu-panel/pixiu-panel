package entity

import "pixiu-panel/pkg/types"

// NotifyLog
// @description: 通知日志
type NotifyLog struct {
	types.BaseDbModel
	UserId  string `gorm:"column:user_id;type:varchar(32);not null;comment:用户id"`
	Pin     string `gorm:"column:pin;type:varchar(32);not null;comment:京东pin"`
	Title   string `gorm:"column:title;type:varchar(255);not null;comment:标题"`
	Content string `gorm:"column:content;type:varchar(500);not null;comment:内容"`
	Status  string `gorm:"column:status;type:varchar(255);not null;comment:推送状态"` // JSON字符串，格式:{"channel": true}
}

// TableName
// @description: 表名
// @receiver NotifyLog
// @return string
func (NotifyLog) TableName() string {
	return "t_notify_log"
}
