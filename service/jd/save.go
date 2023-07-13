package jd

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
	"pixiu-panel/model/param"
	"time"
)

// SaveJdInfo
// @description: 保存京东账户信息
// @param p
// @return err
func SaveJdInfo(p param.SaveJdAccount) (err error) {
	if p.Id == "" {
		// 判断pin是否存在
		var ent entity.UserJd
		_ = db.Client.Take(&ent, "pin = ?", p.Pin).Error
		if ent.Id != "" {
			// 账号已存在，修改更新时间为当前时间
			_ = db.Client.Model(&entity.UserJd{}).
				Where("id = ?", ent.Id).
				Update("last_update", time.Now().Local()).Error
			return
		}

		// 新增
		ent.Pin = p.Pin
		ent.UserId = p.UserId
		err = db.Client.Create(&ent).Error
		return
	}
	// 更新
	up := make(map[string]any)
	up["remark"] = p.Remark
	err = db.Client.Model(&entity.UserJd{}).Where("id = ?", p.Id).Updates(up).Error
	return
}
