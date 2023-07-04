package admin

import (
	"github.com/kataras/iris/v12"
	menuApi "pixiu-panel/api/admin/menu"
)

// menu
// @description: 菜单相关接口
// @param g
func menu(g iris.Party) {
	g.Get("", menuApi.GetMenuTree)
}
