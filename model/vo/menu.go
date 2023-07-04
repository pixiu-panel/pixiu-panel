package vo

// MenuNode
// @description: 菜单节点
type MenuNode struct {
	Path     string     `json:"path"`     // 路径
	Name     string     `json:"name"`     // 名称
	Meta     MenuMeta   `json:"meta"`     // 元数据
	Children []MenuNode `json:"children"` // 子级
}

// MenuMeta
// @description: 菜单元数据
type MenuMeta struct {
	Title      string   `json:"title"`      // 标题
	Icon       string   `json:"icon"`       // 图标
	Rank       int      `json:"rank"`       // 排序(越小越靠前)
	Roles      []string `json:"roles"`      // 页面级别权限设置(传角色code)
	Auths      []string `json:"auths"`      // 按钮级别权限设置(传按钮code)
	ShowParent bool     `json:"showParent"` // 是否显示父级菜单
}
