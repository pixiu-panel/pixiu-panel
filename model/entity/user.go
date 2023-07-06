package entity

import "pixiu-panel/pkg/types"

// User
// @description: 用户表
type User struct {
	types.BaseDbModel
	Username    string          `json:"username" gorm:"index:deleted,unique;type:varchar(255); not null; comment:'登录账号'"`
	Password    string          `json:"password" gorm:"type:varchar(255); comment:'密码'"`
	Email       *string         `json:"email" gorm:"type:varchar(255); comment:'邮箱'"`
	IsVerified  bool            `json:"isVerified" gorm:"type:tinyint(1); not null; default:0; comment:'是否验证邮箱'"`
	Role        string          `json:"role" gorm:"type:enum('admin','user'); default:'user'; not null; comment:'角色'"`
	LastLoginAt *types.DateTime `json:"lastLoginAt" gorm:"comment:'最后登录时间'"`
	LastLoginIp *string         `json:"lastLoginIp" gorm:"type:varchar(255); comment:'最后登录ip'"`
}

// TableName
// @description: 表名
// @receiver User
// @return string
func (User) TableName() string {
	return "t_user"
}
