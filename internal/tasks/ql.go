package tasks

import (
	"gitee.ltd/lxh/logger/log"
	"net/url"
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/ql"
	"pixiu-panel/model/entity"
)

// updateJdAccount
// @description: 更新京东账户信息
func updateJdAccount() {
	// 获取环境变量
	envs := ql.GetEnvs("BBK_V2_WSCK")
	log.Debugf("共获取到 %d 个京东账户", len(envs))
	for _, env := range envs {
		env.Remarks, _ = url.QueryUnescape(env.Remarks)
		pm := make(map[string]any)
		pm["expired"] = env.Status == 1
		pm["last_update"] = env.UpdatedAt
		// 保存京东账户信息
		if err := db.Client.Model(&entity.UserJd{}).Where("pin = ?", env.Remarks).Updates(pm).Error; err != nil {
			log.Errorf("更新京东账户信息失败: %v", err)
		}
	}
}
