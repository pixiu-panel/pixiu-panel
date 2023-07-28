package tasks

import (
	"fmt"
	"gitee.ltd/lxh/logger/log"
	"github.com/duke-git/lancet/v2/slice"
	"net/url"
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/ql"
	"pixiu-panel/internal/ql/model"
	"pixiu-panel/internal/qq"
	"pixiu-panel/internal/wechat"
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
	var expiredPins []string
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

		if data["cookie"].Status == 1 {
			// 过期了
			expiredPins = append(expiredPins, pin)
		}
	}

	// 处理过期通知
	expiredNotify(expiredPins...)
}

// expiredNotify
// @description: 通知账户过期
// @param pin
func expiredNotify(pin ...string) {
	if len(pin) == 0 {
		return
	}

	// 查出所有账号信息
	var jds []entity.UserJd
	if err := db.Client.Find(&jds, "pin IN (?)", pin).Error; err != nil {
		log.Errorf("查询账户信息失败: %v", err)
		return
	}
	// 按用户Id组装成map
	var jdMap = make(map[string][]entity.UserJd)
	var userIds = make([]string, 0)
	for _, jd := range jds {
		jdMap[jd.UserId] = append(jdMap[jd.UserId], jd)
		userIds = append(userIds, jd.UserId)
	}
	// 去重一下用户Id
	userIds = slice.Unique(userIds)
	// 取出所有推送渠道
	var channels []entity.UserNotify
	if err := db.Client.Find(&channels, "user_id IN (?)", userIds).Error; err != nil {
		log.Errorf("查询推送渠道失败: %v", err)
		return
	}
	// 按用户Id组装成map
	var channelMap = make(map[string][]entity.UserNotify)
	for _, channel := range channels {
		channelMap[channel.UserId] = append(channelMap[channel.UserId], channel)
	}

	// 循环过期账号数据，发送推送
	for userId, jdArray := range jdMap {
		var baseMsgArray = []string{"您的京东账号已过期"}
		for _, jd := range jdArray {
			nickname := jd.Nickname
			if jd.Remark != "" {
				nickname = jd.Remark
			}
			baseMsgArray = append(baseMsgArray, fmt.Sprintf("账号: %s\n备注:%s", jd.Pin, nickname))
		}
		// 推送通知
		if css, ok := channelMap[userId]; ok {
			for _, c := range css {
				// 策略发送
				switch c.Channel {
				case "wechat":
					_ = wechat.SendMessage(c.Param, strings.Join(baseMsgArray, "\n"))
				case "qq":
					_ = qq.SendMessage(c.Param, strings.Join(baseMsgArray, "\n"))
				default:
					continue
				}
			}
		}
	}

}
