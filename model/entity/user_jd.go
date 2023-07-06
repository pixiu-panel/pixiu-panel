package entity

import "pixiu-panel/pkg/types"

// UserJd
// @description: 用户京东账号
type UserJd struct {
	types.BaseDbModel
	UserId     string          `json:"userId" gorm:"type:varchar(32);not null;comment:'用户id'"`
	Pin        string          `json:"pin" gorm:"type:varchar(255);not null;comment:'京东pin'"`
	Nickname   string          `json:"nickname" gorm:"type:varchar(255);not null;comment:'京东昵称'"`
	Avatar     string          `json:"avatar" gorm:"type:varchar(255);not null;comment:'京东头像'"`
	Level      string          `json:"level" gorm:"type:varchar(255);not null;comment:'京东等级'"`
	IsPlus     bool            `json:"isPlus" gorm:"type:tinyint(1);not null;default:0;comment:'是否是Plus会员'"`
	Remark     string          `json:"remark" gorm:"type:varchar(255);not null;comment:'备注'"`
	Expired    bool            `json:"expired" gorm:"type:tinyint(1);not null;default:0;comment:'是否过期'"`
	LastUpdate *types.DateTime `json:"lastUpdate" gorm:"comment:'最后更新时间'"`
	QlCookieId int             `json:"qlCookieId" gorm:"type:int(11);not null;comment:'青龙cookie环境变量id'"`
	QlWsckId   int             `json:"qlWsckId" gorm:"type:int(11);not null;comment:'青龙wsck环境变量id'"`
	Cookie     string          `json:"cookie" gorm:"type:varchar(500);not null;comment:'cookie'"`
}

// TableName
// @description: 表名
// @receiver UserJd
// @return string
func (UserJd) TableName() string {
	return "t_user_jd"
}
