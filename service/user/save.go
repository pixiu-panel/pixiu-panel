package user

import (
	"errors"
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/config"
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

// Register
// @description: 注册
// @param p
// @return err
func Register(p param.Register) (err error) {
	// 如果启用了邀请码，校验一下
	if config.Conf.System.Register.InvitationCode {
		log.Debugf("邀请码: %s", p.InvitationCode)
		var count int64
		err = db.Client.Model(&entity.InvitationCode{}).
			Where("code = ?", p.InvitationCode).Where("invitee_id IS NULL").
			Count(&count).Error
		if err != nil {
			err = errors.New("邀请码校验失败")
			return
		}
		if count > 1 {
			err = errors.New("邀请码已被使用")
			return
		}
	}

	// 判断账号是否已存在
	var count int64
	err = db.Client.Model(&entity.User{}).Where("username = ?", p.Username).Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("账号已存在")
		return
	}

	// 加密密码
	utils.PasswordUtils().HashPassword(&p.Password)

	tx := db.Client.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	var ent entity.User
	ent.Username = p.Username
	ent.Password = p.Password
	ent.Role = "user"
	err = tx.Create(&ent).Error
	if err != nil {
		return
	}
	// 修改邀请码数据
	pm := map[string]any{
		"invitee_id": ent.Id,
	}
	err = tx.Model(&entity.InvitationCode{}).Where("code = ?", p.InvitationCode).Updates(pm).Error
	return
}
