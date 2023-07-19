package notify

import (
	"database/sql"
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/internal/db"
	"pixiu-panel/model/entity"
	"regexp"
	"strings"
)

// saveMsgToDb
// @description: 保存消息到数据库
// @param title string 标题
// @param msg []string 消息内容
// @return err error 错误
func saveMsgToDb(title string, msg []string) (err error) {
	var logs []entity.NotifyLog
	for _, m := range msg {
		// 匹配出带账号的行
		accountLine := regexp.MustCompile(`(【)?(京东)?账号\d+(】)?.*`).FindString(m)
		// 去掉前缀
		pin := regexp.MustCompile(`【?(京东)?账号.*】`).ReplaceAllString(accountLine, "")
		// 去掉是否实名信息
		pin = regexp.MustCompile(`\(wskey.*\)?`).ReplaceAllString(pin, "")
		if pin == "" {
			log.Errorf("未找到pin: %s", m)
			continue
		}

		// 清理掉多余的换行
		m = strings.ReplaceAll(m, "\n\n", "")

		// 根据 pin 查出 userId
		var jdAccount entity.UserJd
		err = db.Client.Where("pin = @pin OR nickname = @pin", sql.Named("pin", pin)).
			First(&jdAccount).Error
		if err != nil {
			log.Errorf("[%s]查询用户id失败: %s", pin, err.Error())
			continue
		}

		// 添加记录信息
		logs = append(logs, entity.NotifyLog{
			UserId:  jdAccount.UserId,
			Pin:     jdAccount.Pin,
			Title:   title,
			Content: m,
			Status:  "",
		})
	}

	// 保存记录入库
	err = db.Client.Create(&logs).Error
	// 如果没错，异步发送通知
	if err == nil {
		go sendNotify(logs)
	}
	return
}
