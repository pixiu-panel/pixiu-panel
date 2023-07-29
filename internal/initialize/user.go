package initialize

import (
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
	"pixiu-panel/utils"
)

// initDefaultUser
// @description: 初始化默认用户
func initDefaultUser() {
	// 如果数据库没有有效用户，初始化一个默认账号
	var count int64
	if err := db.Client.Model(entity.User{}).Count(&count).Error; err != nil {
		log.Panicf("初始化默认账号失败: %s", err.Error())
	}
	if count == 0 {
		var ent entity.User
		ent.Username = "pixiu"
		ent.Password = "pixiu123"
		utils.PasswordUtils().HashPassword(&ent.Password)
		ent.Role = "admin"
		if err := db.Client.Create(&ent).Error; err != nil {
			log.Panicf("初始化默认账号失败: %s", err.Error())
		}
	}
}
