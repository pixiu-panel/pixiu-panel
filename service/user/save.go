package user

import (
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
	"time"
)

// UpdateLastLoginInfo
// @description: 更新最后登录时间和IP
// @param id
// @param ip
func UpdateLastLoginInfo(id, ip string) {
	err := db.Client.Model(&entity.User{}).Where("id = ?", id).
		Updates(map[string]any{
			"last_login_at": time.Now(),
			"last_login_ip": ip,
		}).Error
	if err != nil {
		log.Errorf("更新管理员用户最后登录时间和IP失败: %v", err)
	}
}
