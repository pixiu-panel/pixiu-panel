package config

// system
// @description: 系统配置
type system struct {
	Jwt jwt `json:"jwt" yaml:"jwt" mapstructure:"jwt"` // JWT配置
}

// jwt
// @description: JWT配置
type jwt struct {
	Secret string `json:"secret" yaml:"secret" mapstructure:"secret"` // 加密串
}
