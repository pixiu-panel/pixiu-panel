package qq

import (
	"context"
	"fmt"
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"pixiu-panel/internal/qq"
	"pixiu-panel/internal/redis"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"strconv"
	"time"
)

// Notify
// @description: QQ消息回调
// @param ctx
// @return err
func Notify(ctx *gin.Context) {
	// 解析参数
	var bd param.QqCallback
	if err := ctx.ShouldBind(&bd); err != nil {
		log.Errorf("解析QQ消息回调失败: %s", err.Error())
		response.New(ctx).Success()
	}
	// 判断是不是添加好友，如果不是，跳过处理
	if bd.PostType != "request" || bd.RequestType != "friend" {
		response.New(ctx).Success()
		return
	}

	log.Debugf("收到QQ好友添加请求，QQ号: %d  -> %s", bd.UserId, bd.Comment)

	// 组装redisKey
	rdsKey := fmt.Sprintf("notify:bind:waiting:%s", bd.Comment)
	// 如果不存在，直接返回
	if has, _ := redis.Client.Exists(context.Background(), rdsKey).Result(); has == 0 {
		log.Debugf("Redis缓存中不存在该代码: %d", bd.Comment)
		return
	}
	// 存在，同意请求
	if err := qq.AcceptAddFriend(bd.Flag); err != nil {
		// 同意好友请求失败
		log.Errorf("同意好友请求失败: %v", err)
		return
	}
	// 成功，修改Redis缓存数据，设置五分钟内过期
	cacheMsg := fmt.Sprintf("{\"account\": \"%s\",\"nickname\":\"\",\"type\":\"qq\"}", strconv.Itoa(bd.UserId))
	if err := redis.Client.Set(context.Background(), rdsKey, cacheMsg, 5*time.Minute).Err(); err != nil {
		log.Errorf("修改Redis缓存数据失败: %v", err)
	}
}
