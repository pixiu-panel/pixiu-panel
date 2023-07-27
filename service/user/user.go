package user

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/orm"
	"pixiu-panel/model/entity"
	"pixiu-panel/model/param"
	"pixiu-panel/model/vo"
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

// Page
// @description: 分页查询用户列表
// @param p
// @return records
// @return total
// @return err
func Page(p param.PageUser) (records []vo.User, total int64, err error) {
	tx := db.Client.Scopes(orm.Page(p.Current, p.Size)).Table("t_user AS u").
		Joins("LEFT JOIN t_user_jd AS jd ON jd.user_id = u.id").
		Joins("LEFT JOIN t_user_notify AS notify ON notify.user_id = u.id").
		Select("u.*", "COUNT(DISTINCT jd.id) AS bind_jd_count", "COUNT(DISTINCT notify.id) AS bind_notify_count").
		Where("u.is_del = 0").
		Where("jd.is_del = 0").
		Group("u.id").
		Order("u.created_at DESC")

	err = tx.Find(&records).Offset(-1).Limit(-1).Count(&total).Error
	return
}
