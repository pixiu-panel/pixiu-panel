package jd

import (
	"encoding/json"
	"gitee.ltd/lxh/logger/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"pixiu-panel/internal/bbk"
	cacheCli "pixiu-panel/internal/cache"
	"pixiu-panel/pkg/response"
	"pixiu-panel/utils"
	"strings"
)

// GetJdQrcode
// @description: 获取京东二维码
// @param ctx
// @return err
func GetJdQrcode(ctx *gin.Context) {
	// 取出登录用户Id
	userId := ctx.GetString("userId")

	// 生成二维码
	qrcode, err := bbk.GetJdQrcode()
	if err != nil {
		response.New(ctx).SetMsg("获取二维码失败").SetError(err).Fail()
		return
	}

	// 缓存二维码和用户Id关系
	code := strings.ToLower(utils.RandomUtils().GetRandomStringMini(5))
	key := "jd:bind:wait:" + code
	data := map[string]any{
		"userId":  userId,         // 用户Id
		"cookie":  qrcode.Cookie,  // 从BBK获取二维码时的cookie
		"timeout": qrcode.Timeout, // 二维码有效期(秒)
		"status":  0,              // -1已过期 0待扫描 1待确认 2已绑定 3绑定失败
		"pin":     "",             // 绑定成功后返回的pin
	}
	cacheDataBytes, _ := json.Marshal(data)
	// 设置缓存
	err = cacheCli.Cache.Set([]byte(key), cacheDataBytes, qrcode.Timeout)
	if err != nil {
		response.New(ctx).SetMsg("获取二维码失败").Fail()
		return
	}

	// 异步检查绑定状态
	go checkJdBindStatus(key)

	// 替换Cookie为缓存Key
	qrcode.Cookie = code
	// 返回二维码到前端
	response.New(ctx).SetData(qrcode).Success()
}

// CheckBinding
// @description: 检查绑定状态
// @param ctx
func CheckBinding(ctx *gin.Context) {
	// 取出需要校验的Key
	key := ctx.Query("key")
	if key == "" {
		response.New(ctx).SetMsg("参数错误").Fail()
		return
	}
	// 取出缓存的数据
	var che map[string]any
	cbs, err := cacheCli.Cache.Get([]byte("jd:bind:wait:" + key))
	if err != nil {
		log.Errorf("获取缓存失败：%s", err.Error())
		response.New(ctx).SetMsg("二维码已过期").Fail()
		return
	}
	if err = json.Unmarshal(cbs, &che); err != nil {
		log.Errorf("解析缓存失败：%s", err.Error())
		response.New(ctx).SetMsg("参数错误").Fail()
		return
	}
	if ctx.GetString("userId") != che["userId"].(string) {
		response.New(ctx).SetCode(http.StatusForbidden).SetMsg("阁下意欲何为？").Fail()
		return
	}

	// 返回绑定状态
	data := ""
	status := che["status"].(float64)
	switch status {
	// -1已过期 0待扫描 1待确认 2已绑定 3绑定失败
	case -1:
		data = "二维码已失效"
	case 1:
		data = "请在手机上确认登录"
	case 2:
		data = "绑定成功"
	case 3:
		data = "绑定失败"
	}

	response.New(ctx).SetData(data).Success()

}
