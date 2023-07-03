package login

// LoginWithPassword
// @description: 账号密码登录
type LoginWithPassword struct {
	Username string `json:"username" form:"username"` // 用户名
	Password string `json:"password" form:"password"` // 密码
}
