package config

// Conf 全局配置
var Conf config

// config
// @description: 配置
type config struct {
	Db     db     `json:"db" yaml:"db" mapstructure:"db"`             // 数据库配置
	Redis  redis  `json:"redis" yaml:"redis" mapstructure:"redis"`    // Redis配置
	Ql     ql     `json:"ql" yaml:"ql" mapstructure:"ql"`             // 青龙配置
	Notify notify `json:"notify" yaml:"notify" mapstructure:"notify"` // 通知配置
	BBK    bbk    `json:"bbk" yaml:"bbk" mapstructure:"bbk"`          // BBK配置
}
