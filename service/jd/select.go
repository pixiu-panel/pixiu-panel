package jd

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/orm"
	"pixiu-panel/model/entity"
	"pixiu-panel/model/param"
)

// GetBindByUser
// @description: 获取用户绑定的京东账号
// @param p param.PageUserJdAccount 查询参数
// @return records []entity.UserJd 返回数据
// @return total int64 结果总数
// @return err error 错误信息
func GetBindByUser(p param.PageUserJdAccount) (records []entity.UserJd, total int64, err error) {
	err = db.Client.Scopes(orm.Page(p.Current, p.Size)).
		Where("user_id = ?", p.UserId).
		Order("created_at DESC").
		Find(&records).
		Limit(-1).Offset(-1).Count(&total).Error
	return
}
