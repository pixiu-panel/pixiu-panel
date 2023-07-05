package param

// page
// @description: 分页通用参数
type page struct {
	Current int `url:"current,required"` // 页码
	Size    int `url:"size,required"`    // 每页数量
}
