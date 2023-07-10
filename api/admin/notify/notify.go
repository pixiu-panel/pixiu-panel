package notify

import (
	"context"
	"fmt"
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/internal/redis"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/utils"
	"time"
	"unicode/utf8"
)

// Binding
// @description: 绑定推送渠道
// @param ctx
// @return err
func Binding(ctx *gin.Context) {
	var p param.BindingNotify
	if err := ctx.ShouldBind(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(err).Fail()
		return
	}

	// 取出登录用户Id
	userId := ctx.Value("userId").(string)
	log.Debugf("收到绑定推送渠道请求，用户Id：%d", userId)

	// 生成一个Code，作为校验用
	code := utils.RandomUtils().GetRandomStringMini(6)
	rdsKey := fmt.Sprintf("notify:bind:wechat:%s", code)
	// 缓存一个空字符串，表示还没完成校验流程，有效期10分钟
	if err := redis.Client.Set(context.Background(), rdsKey, "", 10*time.Minute).Err(); err != nil {
		log.Errorf("存入Redis失败: %v", err)
		response.New(ctx).SetMsg("绑定失败").SetError(err).Fail()
		return
	}

	// 返回Code
	response.New(ctx).SetData(code).Success()
}

// CheckBinding
// @description: 检查绑定状态
// @param ctx
func CheckBinding(ctx *gin.Context) {
	code := ctx.Query("code")
	if utf8.RuneCountInString(code) != 6 {
		response.New(ctx).SetMsg("参数错误").Fail()
		return
	}
	// 判断key是否存在，不在了就是过期了
	rdsKey := fmt.Sprintf("notify:bind:wechat:%s", code)
	if has, _ := redis.Client.Exists(context.Background(), rdsKey).Result(); has == 0 {
		response.New(ctx).SetMsg("已过期").Fail()
		return
	}
	// 取出值判断是不是空字符串，如果是空的，表示还没完成校验流程，直接返回空数据回去
	if val, _ := redis.Client.Get(context.Background(), rdsKey).Result(); val == "" {
		response.New(ctx).Success()
		return
	} else {
		// 已经有数据了
		response.New(ctx).SetData(val).Success()
	}
}
