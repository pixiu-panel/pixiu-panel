package user

import (
	"errors"
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
	"pixiu-panel/model/param"
	"pixiu-panel/utils"
	"time"
)

// UpdateLastLoginInfo
// @description: 更新最后登录时间和IP
// @param username string 账号
// @param ip string IP地址
func UpdateLastLoginInfo(username, ip string) {
	err := db.Client.Model(&entity.User{}).Where("username = ? OR id = ?", username, username).
		Updates(map[string]any{
			"last_login_at": time.Now(),
			"last_login_ip": ip,
		}).Error
	if err != nil {
		log.Errorf("更新管理员用户最后登录时间和IP失败: %v", err)
	}
}

// ChangePassword
// @description: 修改密码
// @param p
// @return err
func ChangePassword(p param.ChangePassword) (err error) {
	// 先查询老密码
	var dbPassword string
	err = db.Client.Model(&entity.User{}).Select("password").Where("id = ?", p.UserId).
		Pluck("password", &dbPassword).Error
	if err != nil {
		return
	}
	if dbPassword == "" {
		return errors.New("用户不存在")
	}

	// 对比老密码
	if !utils.PasswordUtils().ComparePassword(dbPassword, p.OldPassword) {
		return errors.New("原密码错误")
	}
	// 加密密码
	utils.PasswordUtils().HashPassword(&p.NewPassword)
	// 更新数据
	err = db.Client.Model(&entity.User{}).Where("id = ?", p.UserId).Update("password", p.NewPassword).Error
	return
}
