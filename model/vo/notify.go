package vo

import "pixiu-panel/pkg/types"

// NotifyLog
// @description: 消息通知日志
type NotifyLog struct {
	Id         string         `json:"id"`              // 记录Id
	CreatedAt  types.DateTime `json:"createdAt"`       // 创建时间
	UserId     string         `json:"userId"`          // 用户Id
	Pin        string         `json:"pin"`             // 京东账号
	JdNickname string         `json:"jdNickname"`      // 京东昵称
	Title      string         `json:"title"`           // 标题
	Content    string         `json:"content"`         // 内容
	StatusStr  string         `json:"-"`               // JSON字符串，格式:{"{channel}": true}
	Status     map[string]any `json:"status" gorm:"-"` // 推送状态，格式:{"{channel}": true}
}
