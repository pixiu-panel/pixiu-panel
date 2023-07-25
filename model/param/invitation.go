package param

// PageInvitationCode
// @description: 分页查询邀请码
type PageInvitationCode struct {
	page
	UserId string `json:"userId" form:"userId"` // 用户Id
}
