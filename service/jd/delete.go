package jd

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/ql"
	"pixiu-panel/model/entity"
)

// Delete
// @description: 删除京东账号
// @param userId
// @param id
// @return err
func Delete(userId, id string) (err error) {
	// 开启事务
	tx := db.Client.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 先查出对应的数据，为删除环境变量做准备
	var ent entity.UserJd
	err = tx.Where("user_id = ?", userId).Where("id = ?", id).First(&ent).Error
	if err != nil {
		return
	}

	// 删除数据库内容
	err = tx.Where("user_id = ?", userId).
		Delete(&entity.UserJd{}, "id = ?", id).Error
	if err != nil {
		return
	}
	// 禁用青龙环境变量
	ids := []int{ent.QlCookieId, ent.QlWsckId}
	if err = ql.DisableEnv(ids); err != nil {
		return
	}
	return
}
