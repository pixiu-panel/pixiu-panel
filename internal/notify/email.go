package notify

import (
	"fmt"
	emailService "github.com/jordan-wright/email"
	"net/smtp"
	"pixiu-panel/config"
)

// email
// @description: 邮件通知
type email struct {
	toUser string // 接收人
}

// Send
// @description: 发送通知
// @receiver s
// @param title
// @param content
// @return err
func (s email) Send(title, content string) (err error) {
	// 判断是否启用
	if !config.Conf.Notify.Email.Enable {
		err = fmt.Errorf("邮件通知未启用")
		return
	}
	conf := config.Conf.Notify.Email
	em := emailService.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = fmt.Sprintf("貔貅面板 <%s>", conf.Email)

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{s.toUser}

	// 设置主题
	em.Subject = title

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte(content)

	//设置服务器相关的配置
	auth := smtp.PlainAuth("", conf.Email, conf.Password, conf.Host)
	err = em.Send(conf.Host, auth)
	return
}
