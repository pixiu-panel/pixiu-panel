package wechat

import (
	"encoding/json"
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
)

// Notify
// @description: 微信消息回调
// @param ctx
// @return err
func Notify(ctx *gin.Context) {
	log.Debugf("收到微信消息回调")
	rd, _ := ctx.GetRawData()
	log.Debugf("微信消息回调内容: %s", string(rd))

	// 解析基础结构
	var bd param.WechatCallback
	if err := json.Unmarshal(rd, &bd); err != nil {
		log.Errorf("解析微信消息回调失败: %s", err.Error())
		return
	}

	// 处理消息
	handler(bd)

	response.New(ctx).Success()
}
