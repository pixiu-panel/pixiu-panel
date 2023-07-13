package param

// BindingNotify
// @description: 绑定推送渠道参数
type BindingNotify struct {
	Type string `json:"type" binding:"required"` // 绑定类型 暂时就QQ和微信，后续再完善其他的
}

// NotifyMessage
// @description: 青龙推送消息
type NotifyMessage struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"message" form:"message"`
}

// PageNotifyLog
// @description: 分页获取消息通知日志
type PageNotifyLog struct {
	page
	UserId string `json:"userId" form:"userId"` // 用户id
}
