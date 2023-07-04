package jd

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"pixiu-panel/internal/bbk"
	cacheCli "pixiu-panel/internal/cache"
	"pixiu-panel/model/cache"
	"pixiu-panel/model/param"
	"pixiu-panel/pkg/response"
	"pixiu-panel/pkg/validator"
	"pixiu-panel/service/jd"
	"pixiu-panel/utils"
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
	key := utils.RandomUtils().GetRandomStringMini(5)
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

	// 返回二维码到前端
	response.New(ctx).SetData(qrcode).Success()
}

// GetBind
// @description: 获取用户绑定的京东账号
// @param ctx
func GetBind(ctx iris.Context) {
	var p param.PageUserJdAccount
	if err := ctx.ReadQuery(&p); err != nil {
		response.New(ctx).SetMsg("参数错误").SetError(validator.Translate(err)).Fail()
		return
	}

	// 手动填充用户Id
	p.UserId = ctx.Value("userId").(string)

	// 查询数据
	records, total, err := jd.GetBindByUser(p)
	if err != nil {
		response.New(ctx).SetMsg("获取失败").SetError(err).Fail()
		return
	}
	// 返回数据
	response.New(ctx).SetData(response.NewPageData(records, total, p.Current, p.Size)).Success()
}
