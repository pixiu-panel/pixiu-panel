package vo

import "pixiu-panel/pkg/types"

// InvitationCodeItem
// @description: 邀请码列表数据
type InvitationCodeItem struct {
	Id              string         `json:"id"`              // 邀请码Id
	CreatedAt       types.DateTime `json:"createdAt"`       // 创建时间
	UpdatedAt       types.DateTime `json:"updatedAt"`       // 更新时间
	UserId          string         `json:"userId"`          // 用户Id
	Code            string         `json:"code"`            // 邀请码
	InviteeId       string         `json:"inviteeId"`       // 被邀请人Id
	InviteeUsername string         `json:"inviteeUsername"` // 被邀请人用户名
	InviteeTime     types.DateTime `json:"inviteeTime"`     // 被邀请时间
}
