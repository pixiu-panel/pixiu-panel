package config

import "fmt"

// db
// @description: 数据库配置
type db struct {
	Type     string `json:"type" yaml:"type" mapstructure:"type"`             // 数据库类型
	Host     string `json:"host" yaml:"host" mapstructure:"host"`             // 数据库地址
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`             // 数据库端口
	Username string `json:"username" yaml:"username" mapstructure:"username"` // 数据库用户名
	Password string `json:"password" yaml:"password" mapstructure:"password"` // 数据库密码
	Database string `json:"database" yaml:"database" mapstructure:"database"` // 数据库名称
}

// GetMysqlDSN
// @description: 获取MySQL连接DSN
// @receiver d
// @return string
func (d db) GetMysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Username, d.Password, d.Host, d.Port, d.Database)
}

// GetPostgreSQLDSN
// @description: 获取PostgreSQL连接DSN
// @receiver d
// @return string
func (d db) GetPostgreSQLDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		d.Host, d.Username, d.Password, d.Database, d.Port)
}
