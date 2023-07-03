package jd

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	cacheCli "pixiu-panel/internal/cache"
	"pixiu-panel/model/cache"
	"pixiu-panel/pkg/response"
	"pixiu-panel/utils"
)

// Binding
// @description: 绑定京东账号
// @param ctx
// @return err
func Binding(ctx iris.Context) {
	// 取出登录用户Id
	claims := jwt.Get(ctx).(*cache.JwtCustomClaims)
	userId := claims.Id

	// 生成二维码

	// 缓存二维码和用户Id关系
	key := utils.RandomUtils().GetRandomStringMini(5)
	data := map[string]any{
		"userId": userId, // 用户Id
		"cookie": "",     // 从BBK获取二维码时的cookie
	}
	cacheDataBytes, _ := json.Marshal(data)
	// 10分钟过期
	err := cacheCli.Cache.Set([]byte(key), cacheDataBytes, 60*10)
	if err != nil {
		response.New(ctx).SetMsg("获取二维码失败").Fail()
		return
	}

	// 返回二维码到前端

	return
}
