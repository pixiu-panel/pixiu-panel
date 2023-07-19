package notify

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
)

// DeleteChannel
// @description: 删除已绑定的通知渠道
// @param id
// @return err
func DeleteChannel(id string) (err error) {
	err = db.Client.Delete(&entity.UserNotify{}, "id = ?", id).Error
	return
}
