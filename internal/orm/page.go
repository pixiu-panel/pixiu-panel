package orm

import "gorm.io/gorm"

// Page
// @description: 分页组件
// @param current
// @param size
// @return func(db *gorm.DB) *gorm.DB
func Page(current, size int) func(db *gorm.DB) *gorm.DB {
	// 如果页码是-1，就不分页
	if current == -1 {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
	// 分页
	return func(db *gorm.DB) *gorm.DB {
		if current == 0 {
			current = 1
		}
		if size < 1 {
			size = 10
		}
		// 计算偏移量
		offset := (current - 1) * size
		// 返回组装结果
		return db.Offset(offset).Limit(size)
	}
}
