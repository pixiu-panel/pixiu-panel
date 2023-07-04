package model

import "time"

// Env
// @description: 青龙环境变量
type Env struct {
	Id        int       `json:"id"`        // Id，暂时没啥用
	Value     string    `json:"value"`     // 变量值
	Timestamp string    `json:"timestamp"` // 时间戳
	Status    int       `json:"status"`    // 状态，是否禁用(1应该是禁用)
	Position  int64     `json:"position"`  // 不知道干嘛的
	Name      string    `json:"name"`      // 变量名
	Remarks   string    `json:"remarks"`   // 备注
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}
