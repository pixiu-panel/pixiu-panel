package db

import (
	"gitee.ltd/lxh/logger/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pixiu-panel/config"
)

// Client 客户端
var Client *gorm.DB

// Init
// @description: 初始化数据库连接
func Init() {
	var dialector gorm.Dialector
	switch config.Conf.Db.Type {
	case "mysql":
		// MySQL
		dialector = mysql.Open(config.Conf.Db.GetMysqlDSN())
	case "postgresql":
		// PostgreSQL
		dialector = postgres.Open(config.Conf.Db.GetPostgreSQLDSN())
	default:
		log.Panic("未配置数据库或数据库类型不支持")
	}

	// 初始化连接
	conn, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Panicf("数据库连接初始化失败: %v", err)
	} else {
		log.Debug("数据库连接初始化成功")
	}
	Client = conn
}
