package admin

import (
	"github.com/labstack/echo/v4"
	userApi "pixiu-panel/api/admin/user"
)

// user
// @description: 用户相关接口
// @param g
func user(g *echo.Group) {
	g.GET("", userApi.Info) // 登录用户信息
}
