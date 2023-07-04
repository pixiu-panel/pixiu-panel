package admin

import (
	"github.com/kataras/iris/v12"
	loginApi "pixiu-panel/api/admin/login"
)

// login
// @description: 登录相关接口
// @param g
func login(g iris.Party) {
	g.Post("", loginApi.Login)           // 登录
	g.Post("/refresh", loginApi.Refresh) // 刷新Token
	g.Post("/logout", loginApi.Logout)   // 退出登录
}
