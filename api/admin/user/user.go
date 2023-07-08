package user

import (
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/pkg/response"
)

// Info
// @description: 用户信息
// @param ctx
// @return err
func Info(ctx *gin.Context) {
	log.Debugf("收到获取用户信息请求")

	response.New(ctx).SetData("").Success()
}
