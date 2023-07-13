package notify

import (
	"encoding/json"
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
)

// sendNotify
// @description: 发送通知
func sendNotify(logs []entity.NotifyLog) {
	// 取出所有用户的推送配置
	var notifyConfigs []entity.UserNotify
	err := db.Client.Model(entity.UserNotify{}).Find(&notifyConfigs).Error
	if err != nil {
		log.Errorf("查询用户推送配置失败: %s", err.Error())
		return
	}
	// 梳理成 map
	notifyConfigMap := make(map[string][]entity.UserNotify)
	for _, c := range notifyConfigs {
		notifyConfigMap[c.UserId] = append(notifyConfigMap[c.UserId], c)
	}
	// 挨个推送
	for _, l := range logs {
		log.Debugf("发送通知: %s --> %s", l.Id, l.Pin)
		if configs, ok := notifyConfigMap[l.UserId]; ok {
			pushStatusMap := make(map[string]bool)
			for _, c := range configs {
				log.Debugf("消息 Id: %s --> 推送渠道: %s", l.Id, c.Channel)
				// TODO 推送消息
				pushStatusMap[c.Channel] = true
			}
			bs, _ := json.Marshal(pushStatusMap)
			// 更新推送记录
			db.Client.Model(&entity.NotifyLog{}).Where("id = ?", l.Id).Update("status", string(bs))
		} else {
			log.Debugf("用户 Id: %s 未找到推送配置", l.UserId)
		}
	}
}
