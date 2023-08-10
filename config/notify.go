package config

// notify
// @description: 通知配置
type notify struct {
	AllowTitle []string `json:"allowTitle" yaml:"allowTitle" mapstructure:"allowTitle"` // 允许的标题
	Wechat     wechat   `json:"wechat" yaml:"wechat" mapstructure:"wechat"`             // 微信机器人配置
	QQ         qq       `json:"qq" yaml:"qq" mapstructure:"qq"`                         // QQ机器人(go-cqhttp)配置
	Ftqq       ftqq     `json:"ftqq" yaml:"ftqq" mapstructure:"ftqq"`                   // Server酱配置
	PushDeer   pushDeer `json:"pushDeer" yaml:"pushDeer" mapstructure:"pushDeer"`       // PushDeer配置
	PushPlus   ftqq     `json:"pushPlus" yaml:"pushPlus" mapstructure:"pushPlus"`       // PushPlus配置(复用一下Server酱配置)
	Email      smtp     `json:"smtp" yaml:"smtp" mapstructure:"smtp"`                   // 邮件配置
}

// wechat
// @description: 微信机器人配置
type wechat struct {
	Enable   bool   `json:"enable" yaml:"enable" mapstructure:"enable"`       // 是否启用
	QrCode   string `json:"qrCode" yaml:"qrCode" mapstructure:"qrCode"`       // 二维码解码后的字符串
	Host     string `json:"host" yaml:"host" mapstructure:"host"`             // 微信机器人地址
	Callback string `json:"callback" yaml:"callback" mapstructure:"callback"` // 微信机器人消息回调地址
}

// qq
// @description: QQ机器人(go-cqhttp)配置
type qq struct {
	Enable      bool   `json:"enable" yaml:"enable" mapstructure:"enable"`                // 是否启用
	QrCode      string `json:"qrCode" yaml:"qrCode" mapstructure:"qrCode"`                // 二维码解码后的字符串
	Host        string `json:"host" yaml:"host" mapstructure:"host"`                      // QQ机器人地址
	AccessToken string `json:"accessToken" yaml:"accessToken" mapstructure:"accessToken"` // QQ机器人Token
}

// ftqq
// @description: Server酱配置
type ftqq struct {
	Enable bool `json:"enable" yaml:"enable" mapstructure:"enable"` // 是否启用
}

// pushDeer
// @description: PushDeer配置
type pushDeer struct {
	Enable bool   `json:"enable" yaml:"enable" mapstructure:"enable"` // 是否启用
	Host   string `json:"host" yaml:"host" mapstructure:"host"`       // 自建服务器地址
}

// smtp
// @description: 邮件配置
type smtp struct {
	Enable   bool   `json:"enable" yaml:"enable" mapstructure:"enable"`       // 是否启用
	Host     string `json:"host" yaml:"host" mapstructure:"host"`             // 邮件服务器地址
	Ssl      bool   `json:"ssl" yaml:"ssl" mapstructure:"ssl"`                // 是否启用SSL
	Email    string `json:"email" yaml:"email" mapstructure:"email"`          // 发件邮件地址
	Password string `json:"password" yaml:"password" mapstructure:"password"` // 发件邮件密码
}
