package notify

import (
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
)

// GetNotifyChannel
// @description: 获取用户的通知渠道
// @param userId
// @return records
func GetNotifyChannel(userId string) (records []entity.UserNotify, err error) {
	err = db.Client.Order("created_at DESC").Find(&records, "user_id = ?", userId).Error
	return
}
