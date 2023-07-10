package config

// notify
// @description: 通知配置
type notify struct {
	AllowTitle []string `json:"allowTitle" yaml:"allowTitle" mapstructure:"allowTitle"` // 允许的标题
	Wechat     wechat   `json:"wechat" yaml:"wechat" mapstructure:"wechat"`             // 微信机器人配置
	QQ         qq       `json:"qq" yaml:"qq" mapstructure:"qq"`                         // QQ机器人(go-cqhttp)配置
}

// wechat
// @description: 微信机器人配置
type wechat struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`             // 微信机器人地址
	Callback string `json:"callback" yaml:"callback" mapstructure:"callback"` // 微信机器人消息回调地址
}

// qq
// @description: QQ机器人(go-cqhttp)配置
type qq struct {
	Host        string `json:"host" yaml:"host" mapstructure:"host"`                      // QQ机器人地址
	AccessToken string `json:"accessToken" yaml:"accessToken" mapstructure:"accessToken"` // QQ机器人Token
}
