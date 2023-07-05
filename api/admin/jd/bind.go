package jd

import (
	"encoding/json"
	"gitee.ltd/lxh/logger/log"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"net/http"
	"pixiu-panel/internal/bbk"
	cacheCli "pixiu-panel/internal/cache"
	"pixiu-panel/model/cache"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/service/jd"
	"pixiu-panel/utils"
	"regexp"
)

// GetJdQrcode
// @description: 获取京东二维码
// @param ctx
// @return err
func GetJdQrcode(ctx iris.Context) {
	// 取出登录用户Id
	claims := jwt.Get(ctx).(*cache.JwtCustomClaims)
	userId := claims.Id

	// 生成二维码
	qrcode, err := bbk.GetJdQrcode()
	if err != nil {
		response.New(ctx).SetMsg("获取二维码失败").SetError(err).Fail()
		return
	}

	// 缓存二维码和用户Id关系
	key := "jd:bind:wait:" + utils.RandomUtils().GetRandomStringMini(5)
	data := map[string]any{
		"userId": userId,        // 用户Id
		"cookie": qrcode.Cookie, // 从BBK获取二维码时的cookie
	}
	cacheDataBytes, _ := json.Marshal(data)
	// 设置缓存
	err = cacheCli.Cache.Set([]byte(key), cacheDataBytes, qrcode.Timeout)
	if err != nil {
		response.New(ctx).SetMsg("获取二维码失败").Fail()
		return
	}

	// 替换Cookie为缓存Key
	qrcode.Cookie = key
	// 返回二维码到前端
	response.New(ctx).SetData(qrcode).Success()
}

// CheckBinding
// @description: 检查绑定状态
// @param ctx
func CheckBinding(ctx iris.Context) {
	// 取出需要校验的Key
	key := ctx.URLParam("key")
	if key == "" {
		response.New(ctx).SetMsg("参数错误").Fail()
		return
	}
	// 取出缓存的数据
	var che map[string]string
	cbs, err := cacheCli.Cache.Get([]byte("jd:bind:wait:" + key))
	if err != nil {
		log.Errorf("获取缓存失败：%s", err.Error())
		response.New(ctx).SetMsg("参数错误").Fail()
		return
	}
	if err = json.Unmarshal(cbs, &che); err != nil {
		log.Errorf("解析缓存失败：%s", err.Error())
		response.New(ctx).SetMsg("参数错误").Fail()
		return
	}
	if ctx.Value("userId").(string) != che["userId"] {
		response.New(ctx).SetCode(http.StatusForbidden).SetMsg("阁下意欲何为？").Fail()
		return
	}

	// 查询二维码扫描状态
	status, err := bbk.CheckJdQrcode(che["cookie"])
	if err != nil {
		response.New(ctx).SetMsg("获取绑定状态失败").SetError(err).Fail()
		return
	}
	switch status.Code {
	case 200:
		// 还没扫描，直接返回空数据回去
		response.New(ctx).Success()
		return
	case 500, 202, 408:
		// 二维码失效
		response.New(ctx).SetMsg("二维码已失效").Success()
		return
	case 201:
		// 请在手机上确认登录
		response.New(ctx).SetMsg("请在手机上确认登录").Success()
		return
	case 410:
		// 登录成功
		// 提取出用户的PIN，准备入库
		pinMatch := regexp.MustCompile(`\[(.*?)\]`).FindStringSubmatch(status.Data.Msg)
		if len(pinMatch) != 2 {
			response.New(ctx).SetMsg("绑定失败，未获取到pin").Fail()
			return
		}
		pin := pinMatch[1]
		// 入库
		pm := param.SaveJdAccount{
			Pin:    pin,
			UserId: che["userId"],
		}
		if err = jd.SaveJdInfo(pm); err != nil {
			log.Errorf("保存京东账号失败: %v", err)
			response.New(ctx).SetMsg("绑定失败").SetError(err).Fail()
		}
		return
	}
	response.New(ctx).Success()

}
