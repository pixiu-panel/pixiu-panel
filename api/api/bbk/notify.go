package bbk

import (
	"errors"
	"fmt"
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/url"
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/notify"
	"pixiu-panel/model/entity"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"time"
)

// Notify
//
//	@Description: BBK回调
//	@param ctx
func Notify(ctx *gin.Context) {
	code := ctx.Query("code")
	log.Debugf("BBK回调类型: %s", code)

	var p param.BBKNotify
	if err := ctx.ShouldBind(&p); err != nil {
		log.Debugf("参数错误: %s", err.Error())
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}
	// URL 编码处理一下，防止中文乱码
	p.Pin, _ = url.QueryUnescape(p.Pin)

	dealScanAfter(p)
	response.New(ctx).Success()
}

// dealScanAfter
//
//	@Description: 处理扫码后逻辑
//	@param pin
func dealScanAfter(p param.BBKNotify) {
	// 查询Pin是否已经入库，如果没入库，休眠一秒钟，然后递归
	var jdInfo entity.UserJd
	err := db.Client.Model(entity.UserJd{}).Where("pin = ?", p.Pin).First(&jdInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 休眠一秒钟后再查，有可能是取BBK扫码结果的任务还没做完
			time.Sleep(time.Second)
			dealScanAfter(p)
		} else {
			log.Errorf("查询用户京东信息失败: %s", err.Error())
			return
		}
	}

	log.Debugf("京东账号: %s", jdInfo.Nickname)

	// 修改状态为正常
	err = db.Client.Model(&entity.UserJd{}).Where("pin = ?", p.Pin).Update("expired", false).Error
	if err != nil {
		log.Errorf("更新用户京东状态失败: %s", err.Error())
		return
	}

	// 获取绑定的通知渠道
	var notifyConfigs []entity.UserNotify
	err = db.Client.Model(entity.UserNotify{}).Where("user_id = ?", jdInfo.UserId).Find(&notifyConfigs).Error
	if err != nil {
		log.Errorf("查询用户推送配置失败: %s", err.Error())
		return
	}
	// 组装消息内容
	msg := fmt.Sprintf("账号: %s\n昵称: %s\n状态: 成功", jdInfo.Pin, jdInfo.Nickname)
	// 循环发送
	for _, c := range notifyConfigs {
		_ = notify.New(c.Channel, c.Param).Send(p.Title, msg)
	}
}
