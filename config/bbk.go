package config

// bbk
// @description: BBK配置
type bbk struct {
	JdQr     jdQr     `json:"jdQr" yaml:"jdQr" mapstructure:"jdQr"`             // 京东扫码服务配置
	WechatQr wechatQr `json:"wechatQr" yaml:"wechatQr" mapstructure:"wechatQr"` // 微信扫码服务配置
	Sms      sms      `json:"sms" yaml:"sms" mapstructure:"sms"`                // 短信登录配置
}

// jdQr
// @description: 京东扫码服务配置
type jdQr struct {
	Enable bool   `json:"enable" yaml:"enable" mapstructure:"enable"` // 是否启用京东扫码服务
	Host   string `json:"host" yaml:"host" mapstructure:"host"`       // 京东扫码服务地址
}

// wechatQr
// @description: 微信扫码服务配置
type wechatQr struct {
	Enable bool   `json:"enable" yaml:"enable" mapstructure:"enable"` // 是否启用微信扫码服务
	Host   string `json:"host" yaml:"host" mapstructure:"host"`       // 微信扫码服务地址
}

// sms
// @description: 短信登录配置
type sms struct {
	Enable bool   `json:"enable" yaml:"enable" mapstructure:"enable"` // 是否启用短信登录服务
	Host   string `json:"host" yaml:"host" mapstructure:"host"`       // 短信登录服务地址
}
