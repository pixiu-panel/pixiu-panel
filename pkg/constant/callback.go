package constant

// WechatMessageType 微信消息类型
type WechatMessageType int

const (
	WechatMessageTypeText      WechatMessageType = 1  // 文字消息
	WechatMessageTypeAddFriend WechatMessageType = 37 // 添加好友
)
