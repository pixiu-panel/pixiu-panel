package config

import "fmt"

// Redis配置
type redis struct {
	Host     string `mapstructure:"host" yaml:"host"`         // 主机
	Port     int    `mapstructure:"port" yaml:"port"`         // 端口
	Password string `mapstructure:"password" yaml:"password"` // 密码
	Db       int    `mapstructure:"db" yaml:"db"`             // 数据库名称
}

func (r redis) GetDSN() string {
	return fmt.Sprintf("%s:%v", r.Host, r.Port)
}
