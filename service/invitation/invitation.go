package invitation

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/orm"
	"pixiu-panel/model/param"
	"pixiu-panel/model/vo"
)

// Page
// @description: 分页查询邀请码
// @param p
// @return records
// @return total
// @return err
func Page(p param.PageInvitationCode) (records []vo.InvitationCodeItem, total int64, err error) {
	tx := db.Client.Scopes(orm.Page(p.Current, p.Size)).
		Table("t_invitation_code AS tic").
		Joins("LEFT JOIN t_user AS tu ON tic.invitee_id = tu.id").
		Select("tic.*", "tu.username AS invitee_username", "tu.created_at AS invitee_time")
	if p.UserId != "" {
		tx.Where("tic.user_id = ?", p.UserId)
	}
	err = tx.Find(&records).Offset(-1).Limit(-1).Count(&total).Error
	return
}
