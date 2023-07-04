package initialize

import "pixiu-panel/pkg/auth"

// InitSystem
// @description: 初始化系统
func InitSystem() {
	initConfig()            // 初始化配置
	databaseTable()         // 初始化数据库表
	auth.InitOAuth2Server() // 初始化OAuth2服务
}
