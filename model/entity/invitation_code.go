package entity

import "pixiu-panel/pkg/types"

// InvitationCode
// @description: 邀请码表
type InvitationCode struct {
	types.BaseDbModel
	UserId    string  ` json:"userId" gorm:"type:varchar(32);not null;comment:用户Id"`
	Code      string  ` json:"code" gorm:"type:varchar(8);not null;comment:邀请码"`
	InviteeId *string ` json:"inviteeId" gorm:"type:varchar(32);comment:被邀请人Id"`
}

// TableName
// @description: 表名
// @receiver InvitationCode
// @return string
func (InvitationCode) TableName() string {
	return "t_invitation_code"
}
