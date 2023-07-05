package param

// PageUserJdAccount
// @description: 分页查询用户京东账号
type PageUserJdAccount struct {
	page
	UserId string // 用户Id，手动填充
}

// SaveJdAccount
// @description: 保存京东账户信息
type SaveJdAccount struct {
	Id     string `json:"id" form:"id"`
	Pin    string // 用户pin - 新增时手动填充
	UserId string // 用户Id - 新增时手动填充
	Remark string `json:"remark" form:"remark"` // 备注
}
