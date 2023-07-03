package user

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
)

// GetUserWithLogin
// @description: 登录时获取用户信息
// @param keyword
// @return ent
// @return err
func GetUserWithLogin(keyword string) (ent entity.User, err error) {
	err = db.Client.Where("username = ? OR email = ?", keyword, keyword).First(&ent).Error
	return
}
