package login

// LoginWithPassword
// @description: 账号密码登录
type LoginWithPassword struct {
	Username string `json:"username" form:"username" validate:"required"` // 用户名
	Password string `json:"password" form:"password" validate:"required"` // 密码
}
