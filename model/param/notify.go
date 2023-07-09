package param

// BindingNotify
// @description: 绑定推送渠道参数
type BindingNotify struct {
	Type string `json:"type" validate:"required"`
}

// NotifyMessage
// @description: 推送消息
type NotifyMessage struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"message" form:"message"`
}
