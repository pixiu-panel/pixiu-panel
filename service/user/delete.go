package user

import (
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/ql"
	"pixiu-panel/model/entity"
)

// Delete
// @description: 删除用户
// @param id
// @return err
func Delete(id string) (err error) {
	tx := db.Client.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 删除账号信息
	if err = tx.Delete(&entity.User{}, "id = ?", id).Error; err != nil {
		log.Errorf("删除用户账号信息失败: %v", err)
		return
	}

	// 查询已绑定的京东账号
	var jds []entity.UserJd
	if err = tx.Find(&jds, "user_id = ?", id).Error; err != nil {
		log.Errorf("查询已绑定的京东账号失败: %v", err)
		return
	}
	var envIds []int
	for _, jd := range jds {
		envIds = append(envIds, jd.QlCookieId)
		envIds = append(envIds, jd.QlWsckId)
	}
	// 查询推送渠道
	var notifies []entity.UserNotify
	if err = tx.Find(&notifies, "user_id = ?", id).Error; err != nil {
		log.Errorf("查询已绑定的推送渠道失败: %v", err)
		return
	}

	// 删除推送渠道
	if err = tx.Delete(&entity.UserNotify{}, "user_id = ?", id).Error; err != nil {
		log.Errorf("删除推送渠道失败: %v", err)
		return
	}
	// TODO 删除微信好友
	// 这个未来再做，因为删除微信好友有bug，删除不彻底

	// 删除京东账号
	if err = tx.Delete(&entity.UserJd{}, "user_id = ?", id).Error; err != nil {
		log.Errorf("删除京东账号失败: %v", err)
		return
	}
	// 禁用青龙环境变量
	if err = ql.DisableEnv(envIds); err != nil {
		log.Errorf("禁用青龙环境变量失败: %v", err)
		return
	}
	return
}
