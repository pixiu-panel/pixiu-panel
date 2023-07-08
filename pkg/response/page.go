package response

// PageData
// @description: 分页数据通用结构体
type PageData[T any] struct {
	Current   int   `json:"current"`   // 当前页码
	Size      int   `json:"size"`      // 每页数量
	Total     int64 `json:"total"`     // 总数
	TotalPage int   `json:"totalPage"` // 总页数
	Records   T     `json:"records"`   // 返回数据
}

// NewPageData
// @description: 创建分页数据
// @param records any 数据列表
// @param total int64 总数
// @param current int 页码
// @param size int 页数量
// @return data PageData[any] 分页数据
func NewPageData(records any, total int64, current, size int) (data PageData[any]) {
	// 处理一下页码、页数量
	if current == -1 {
		current = 1
		size = int(total)
	}
	// 计算总页码
	totalPage := 0
	if total > 0 {
		upPage := 0
		if int(total)%size > 0 {
			upPage = 1
		}
		totalPage = (int(total) / size) + upPage
	}
	data = PageData[any]{
		Current:   current,
		Size:      size,
		Total:     total,
		TotalPage: totalPage,
		Records:   records,
	}
	return
}
