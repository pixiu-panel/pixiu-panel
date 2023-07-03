package admin

import (
	"github.com/labstack/echo/v4"
	loginApi "pixiu-panel/api/admin/login"
)

// login
// @description: 登录相关接口
// @param g
func login(g *echo.Group) {
	g.POST("", loginApi.Login)
}
