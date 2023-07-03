package admin

import (
	"github.com/kataras/iris/v12"
	jdApi "pixiu-panel/api/admin/jd"
)

// jd
// @description: 京东相关接口
// @param g
func jd(g iris.Party) {
	bind := g.Party("/binding")
	bind.Get("/qrcode", jdApi.GetJdQrcode) // 获取京东二维码
}
