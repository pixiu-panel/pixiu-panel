package param

// page
// @description: 分页通用参数
type page struct {
	Current int `json:"current" form:"current"  url:"current" binding:"required"` // 页码
	Size    int `json:"size" form:"size" url:"size" binding:"required"`           // 每页数量
}
