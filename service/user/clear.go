package user

import (
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
)

// ClearEmptyUser
// @description: 清理空用户
func ClearEmptyUser() {
	subQueryJd := db.Client.Model(&entity.UserJd{}).Select("user_id")
	subQueryNotify := db.Client.Model(&entity.UserNotify{}).Select("user_id")

	err := db.Client.Where("id NOT IN (?)", subQueryJd).
		Where("id NOT IN (?)", subQueryNotify).
		Where("`role` != 'admin'").
		Delete(&entity.User{}).Error
	if err != nil {
		log.Errorf("清理空用户失败: %v", err)
	}
}
