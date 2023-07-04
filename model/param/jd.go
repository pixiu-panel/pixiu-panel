package param

// PageUserJdAccount
// @description: 分页查询用户京东账号
type PageUserJdAccount struct {
	page
	UserId string `json:"-" form:"-"` // 用户Id，手动填充
}

// SaveJdAccount
// @description: 保存京东账户信息
type SaveJdAccount struct {
	Id     string `json:"id" form:"id"`
	Remark string `json:"remark" form:"remark"` // 备注
}
