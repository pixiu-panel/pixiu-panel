package config

// ql
// @description: 青龙配置
type ql struct {
	Host         string `json:"host" yaml:"host" mapstructure:"host"`                         // 青龙面板地址
	ClientId     string `json:"clientId" yaml:"clientId" mapstructure:"clientId"`             // 青龙面板ClientId
	ClientSecret string `json:"clientSecret" yaml:"clientSecret" mapstructure:"clientSecret"` // 青龙面板ClientSecret
}
