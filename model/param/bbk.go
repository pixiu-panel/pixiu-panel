package param

// BBKNotify
// @Description: BBK通知消息
type BBKNotify struct {
	Title string `json:"title" form:"title"` // 通知类型
	Pin   string `json:"pin" form:"pin"`     // PIN
}
