package notify

// BindingNotify
// @description: 绑定推送渠道参数
type BindingNotify struct {
	Type string `json:"type" validate:"required"`
}
