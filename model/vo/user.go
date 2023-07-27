package vo

import "pixiu-panel/pkg/types"

// User
// @description: 用户列表
type User struct {
	Id              string          `json:"id"`              // 用户Id
	Username        string          `json:"username"`        // 用户名
	Email           string          `json:"email"`           // 邮箱
	IsVerified      bool            `json:"isVerified"`      // 是否验证邮箱
	Role            string          `json:"role"`            // 角色
	CreatedAt       types.DateTime  `json:"createdAt"`       // 注册时间
	LastLoginAt     *types.DateTime `json:"lastLoginAt"`     // 最后登录时间
	LastLoginIp     *string         `json:"lastLoginIp"`     // 最后登录ip
	BindJdCount     int64           `json:"bindJdCount"`     // 绑定的京东账号数量
	BindNotifyCount int64           `json:"bindNotifyCount"` // 绑定的通知渠道数量
}
