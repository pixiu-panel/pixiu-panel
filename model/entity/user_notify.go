package entity

import "pixiu-panel/pkg/types"

// UserNotify
// @description: 用户通知配置表
type UserNotify struct {
	types.BaseDbModelWithReal
	UserId  string `json:"userId" gorm:"type:varchar(32); not null; comment:'用户id'"`
	Channel string `json:"channel" gorm:"type:varchar(32); not null; comment:'通知渠道'"`
	Param   string `json:"param" gorm:"type:varchar(500); not null; comment:'通知参数'"` // 暂时就用来存微信Id或者QQ号
}

// TableName
// @description: 表名
// @receiver UserNotify
// @return string
func (UserNotify) TableName() string {
	return "t_user_notify"
}
