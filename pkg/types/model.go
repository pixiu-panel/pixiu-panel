package types

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"strings"
)

// BaseDbModel
// @description: 带逻辑删除的数据库通用字段
type BaseDbModel struct {
	Id        string                `json:"id" gorm:"type:varchar(32);primarykey"`
	CreatedAt DateTime              `json:"createdAt"`
	UpdatedAt DateTime              `json:"updatedAt"`
	DeletedAt int64                 `json:"-" gorm:"index:deleted; default:0"`
	IsDel     soft_delete.DeletedAt `json:"-" gorm:"softDelete:flag,DeletedAtField:DeletedAt; index:deleted; default:0; type:tinyint(1)"`
}

// BeforeCreate 创建数据库对象之前生成UUID
func (m *BaseDbModel) BeforeCreate(*gorm.DB) (err error) {
	if m.Id == "" {
		m.Id = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}

// =====================================================================================================================

// BaseDbModelWithReal
// @description: 不带逻辑删除的数据库通用字段
type BaseDbModelWithReal struct {
	Id        string   `json:"id" gorm:"type:varchar(32);primarykey"`
	CreatedAt DateTime `json:"createdAt"`
	UpdatedAt DateTime `json:"updatedAt"`
}

// BeforeCreate 创建数据库对象之前生成UUID
func (m *BaseDbModelWithReal) BeforeCreate(*gorm.DB) (err error) {
	m.Id = strings.ReplaceAll(uuid.New().String(), "-", "")
	return
}
