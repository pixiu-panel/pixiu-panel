package param

// PageUser
// @description: 分页查询用户
type PageUser struct {
	page
}

// ChangePassword
// @description:
type ChangePassword struct {
	UserId          string `json:"userId" form:"userId"`                                                            // 用户ID
	OldPassword     string `json:"oldPassword" form:"oldPassword" binding:"required"`                               // 旧密码
	NewPassword     string `json:"newPassword" form:"newPassword" binding:"required"`                               // 新密码
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" binding:"required,eqcsfield=NewPassword"` // 确认新密码
}
