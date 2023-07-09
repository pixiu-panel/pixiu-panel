package notify

import (
	"gitee.ltd/lxh/logger/log"
	"pixiu-panel/model/param"
	"regexp"
)

// Parse
// @description: 解析消息
// @param msg
// @return err
func Parse(msg param.NotifyMessage) (err error) {
	log.Debugf("待解析消息: \n%s", msg.Content)

	// 去掉不需要的内容
	// 去掉尾巴上的"本通知"
	content := regexp.MustCompile(`\n+本通知.*`).ReplaceAllString(msg.Content, "")
	// 去掉尾巴上的"入口"
	content = regexp.MustCompile(`\n+入口.*`).ReplaceAllString(content, "")

	// 提取内容正则
	re := regexp.MustCompile(`(【)?(京东)?账号\d+(】)?.*?`)
	matches := re.FindAllStringSubmatchIndex(content, -1)

	msgArr := make([]string, 0, len(matches))
	for i, match := range matches {
		if i == len(matches)-1 {
			msgArr = append(msgArr, content[match[0]:])
		} else {
			msgArr = append(msgArr, content[match[0]:matches[i+1][0]])
		}
	}
	log.Debugf("共有%d条待处理消息", len(msgArr))

	// 入库
	err = saveMsgToDb(msg.Title, msgArr)
	return
}
