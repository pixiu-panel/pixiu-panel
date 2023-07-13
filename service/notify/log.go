package notify

import (
	"encoding/json"
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/orm"
	"pixiu-panel/model/param"
	"pixiu-panel/model/vo"
)

// PageNotifyLog
// @description: 分页获取消息通知日志
// @param p
// @return records
// @return total
// @return err
func PageNotifyLog(p param.PageNotifyLog) (records []vo.NotifyLog, total int64, err error) {
	tx := db.Client.Scopes(orm.Page(p.Current, p.Size)).
		Table("t_notify_log AS tnl").
		Joins("LEFT JOIN t_user_jd tuj ON tnl.pin = tuj.pin").
		Select("tnl.id", "tnl.created_at", "tnl.title", "tnl.content",
			"tnl.`status` AS status_str", "tnl.pin", "tuj.nickname AS jd_nickname").
		Order("tnl.created_at DESC")

	// 指定用户
	if p.UserId != "" {
		tx.Where("tnl.user_id = ?", p.UserId)
	}

	err = tx.Find(&records).Offset(-1).Limit(-1).Count(&total).Error
	if err == nil {
		for i, record := range records {
			if record.StatusStr != "" {
				_ = json.Unmarshal([]byte(record.StatusStr), &records[i].Status)
			}
		}
	}
	return
}
