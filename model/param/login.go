package param

// LoginWithPassword
// @description: 账号密码登录
type LoginWithPassword struct {
	Username string `json:"username" form:"username" binding:"required"` // 用户名
	Password string `json:"password" form:"password" binding:"required"` // 密码
}

// RefreshToken
// @description: 刷新Token入参
type RefreshToken struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"` // 刷新Token
	GrantType    string `json:"grant_type" form:"grant_type" binding:"required"`       // 授权类型,写refresh_token
}

// Register
// @description: 注册账号
type Register struct {
	Username       string `json:"username" form:"username" binding:"required"` // 用户名
	Password       string `json:"password" form:"password" binding:"required"` // 密码
	InvitationCode string `json:"invitationCode" form:"invitationCode"`        // 邀请码
}
