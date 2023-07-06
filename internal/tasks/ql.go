package tasks

import (
	"gitee.ltd/lxh/logger/log"
	"net/url"
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/ql"
	"pixiu-panel/internal/ql/model"
	"pixiu-panel/model/entity"
	"strings"
)

// updateJdAccount
// @description: 更新京东账户信息
func updateJdAccount() {
	// 获取环境变量
	envs := ql.GetEnvs("")

	// 循环处理一下，把数据解析为map，{"pin": {"wsck": ent, "cookie": ent}}
	var envMap = make(map[string]map[string]model.Env)
	for _, env := range envs {
		switch env.Name {
		case "BBK_V2_WSCK":
			// 是wsck
			pin, _ := url.QueryUnescape(env.Remarks)
			if _, ok := envMap[pin]; !ok {
				// 不存在，初始化一下
				envMap[pin] = make(map[string]model.Env)
			}
			envMap[pin]["wsck"] = env
		case "JD_COOKIE":
			// 是cookie
			// 解析一下，取出pin
			cookie := env.Value
			cookie, _ = url.QueryUnescape(cookie)
			// 先按;分割，再按=分割，取出pt_pin
			for _, s := range strings.Split(cookie, ";") {
				if strings.Contains(s, "pt_pin") {
					// 是pt_pin
					pin := strings.Split(s, "=")[1]
					log.Debugf("获取到京东账户: %s", pin)
					if _, ok := envMap[pin]; !ok {
						// 不存在，初始化一下
						envMap[pin] = make(map[string]model.Env)
					}
					env.Remarks = pin
					envMap[pin]["cookie"] = env
				}
			}
		default:
			// 其他的不处理
		}
	}

	log.Debugf("共获取到 %d 个京东账户", len(envMap))
	for pin, data := range envMap {
		pm := make(map[string]any)
		pm["expired"] = data["cookie"].Status == 1
		pm["last_update"] = data["wsck"].UpdatedAt
		pm["cookie"] = data["cookie"].Value
		pm["ql_cookie_id"] = data["cookie"].Id
		pm["ql_wsck_id"] = data["wsck"].Id

		// 保存京东账户信息
		if err := db.Client.Model(&entity.UserJd{}).Where("pin = ?", pin).Updates(pm).Error; err != nil {
			log.Errorf("更新京东账户信息失败: %v", err)
		}
	}
}
