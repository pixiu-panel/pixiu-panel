package config

// system
// @description: 系统配置
type system struct {
	Domain   string   `json:"domain" yaml:"domain" mapstructure:"domain"`       // 绑定域名
	Register register `json:"register" yaml:"register" mapstructure:"register"` // 注册配置
}

type register struct {
	Enable         bool `json:"enable" yaml:"enable" mapstructure:"enable"`                         // 是否开启注册
	InvitationCode bool `json:"invitationCode" yaml:"invitationCode" mapstructure:"invitationCode"` // 是否开启邀请码
}
