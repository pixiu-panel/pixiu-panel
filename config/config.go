package config

// Conf 全局配置
var Conf config

// config
// @description: 配置
type config struct {
	System system `json:"system" yaml:"system" mapstructure:"system"` // 系统配置
	Ql     ql     `json:"ql" yaml:"ql" mapstructure:"ql"`             // 青龙配置
	Notify notify `json:"notify" yaml:"notify" mapstructure:"notify"` // 通知配置
	BBK    bbk    `json:"bbk" yaml:"bbk" mapstructure:"bbk"`          // BBK配置
}
