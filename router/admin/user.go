package admin

import (
	"github.com/kataras/iris/v12"
	userApi "pixiu-panel/api/admin/user"
)

// user
// @description: 用户相关接口
// @param g
func user(g iris.Party) {
	g.Get("", userApi.Info) // 登录用户信息
}
