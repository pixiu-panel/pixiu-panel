package entity

import "pixiu-panel/pkg/types"

// UserJd
// @description: 用户京东账号
type UserJd struct {
	types.BaseDbModel
	UserId     string          `json:"userId" gorm:"type:varchar(32);not null;comment:'用户id'"`
	Pin        string          `json:"pin" gorm:"type:varchar(255);not null;comment:'京东pin'"`
	Avatar     string          `json:"avatar" gorm:"type:varchar(255);not null;comment:'京东头像'"`
	Remark     string          `json:"remark" gorm:"type:varchar(255);not null;comment:'备注'"`
	Expired    bool            `json:"expired" gorm:"type:tinyint(1);not null;default:0;comment:'是否过期'"`
	LastUpdate *types.DateTime `json:"lastUpdate" gorm:"comment:'最后更新时间'"`
}

// TableName
// @description: 表名
// @receiver UserJd
// @return string
func (UserJd) TableName() string {
	return "t_user_jd"
}
