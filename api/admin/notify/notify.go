package notify

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/internal/db"
	"pixiu-panel/internal/redis"
	"pixiu-panel/model/entity"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/notify"
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
	userId := ctx.GetString("userId")
	log.Debugf("收到绑定推送渠道请求，用户Id：%s", userId)

	// 生成一个Code，作为校验用
	code := utils.RandomUtils().GetRandomStringMini(6)
	rdsKey := fmt.Sprintf("notify:bind:waiting:%s", code)
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
	rdsKey := fmt.Sprintf("notify:bind:waiting:%s", code)
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
		var ui map[string]string
		if err := json.Unmarshal([]byte(val), &ui); err != nil {
			log.Errorf("解析Redis数据失败: %v", err)
			response.New(ctx).SetMsg("绑定失败").SetError(err).Fail()
			return
		}

		// 取出用户Id
		userId := ctx.GetString("userId")
		// 保存数据入库
		var ent entity.UserNotify
		ent.UserId = userId
		ent.Channel = ui["type"]
		ent.Param = ui["account"]
		if err := db.Client.Create(&ent).Error; err != nil {
			log.Errorf("保存数据失败: %v", err)
			response.New(ctx).SetMsg("绑定失败").SetError(err).Fail()
			return
		}

		response.New(ctx).SetData(ui).Success()
	}
}

// GetBindingAccounts
// @description: 获取绑定的推送渠道
// @param ctx
func GetBindingAccounts(ctx *gin.Context) {
	records, err := notify.GetNotifyChannel(ctx.GetString("userId"))
	if err != nil {
		log.Errorf("获取绑定的推送渠道失败: %v", err)
		response.New(ctx).SetMsg("获取绑定的推送渠道失败").SetError(err).Fail()
		return
	}
	response.New(ctx).SetData(records).Success()
}
