package jd

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
	"pixiu-panel/model/param"
)

// SaveJdInfo
// @description: 保存京东账户信息
// @param p
// @return err
func SaveJdInfo(p param.SaveJdAccount) (err error) {
	if p.Id == "" {
		// 新增
		jd := entity.UserJd{
			Pin:    p.Pin,
			UserId: p.UserId,
		}
		err = db.Client.Create(&jd).Error
		return
	}
	// 更新
	up := make(map[string]any)
	up["remark"] = p.Remark
	err = db.Client.Model(&entity.UserJd{}).Where("id = ?", p.Id).Updates(up).Error
	return
}
