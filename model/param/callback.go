package param

import "pixiu-panel/pkg/constant"

// WechatCallback
// @description: 微信消息回调
type WechatCallback struct {
	Content   string                     `json:"content"`   // 消息正文
	FromGroup string                     `json:"fromGroup"` // 来源群组
	FromUser  string                     `json:"fromUser"`  // 来源用户
	Time      string                     `json:"time"`      // 消息时间
	Timestamp int                        `json:"timestamp"` // 消息时间戳
	Type      constant.WechatMessageType `json:"type"`
}

// WechatAddFriend
// @description: 添加好友消息
type WechatAddFriend struct {
	FromUsername    string `xml:"fromusername,attr"`    // 来源用户微信Id
	FromNickname    string `xml:"fromnickname,attr"`    // 来源用户昵称
	Content         string `xml:"content,attr"`         // 附加消息内容
	EncryptUsername string `xml:"encryptusername,attr"` // V3数据，同意添加好友需要
	Ticket          string `xml:"ticket,attr"`          // V4数据，同意添加好友需要
}
