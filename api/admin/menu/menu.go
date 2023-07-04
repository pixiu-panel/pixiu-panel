package menu

import (
	"github.com/kataras/iris/v12"
	"pixiu-panel/model/vo"
	"pixiu-panel/pkg/response"
)

// GetMenuTree
// @description: 获取菜单树
// @param ctx
func GetMenuTree(ctx iris.Context) {
	menus := make([]vo.MenuNode, 0)
	response.New(ctx).SetData(menus).Success()
}
