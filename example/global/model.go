package global

import (
	"gorm.io/gorm"
	"time"
)

// 公共模型
type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey"` // 主键ID
	CreatedAt time.Time      `json:"created_at"`           // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`           // 创建时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`       // 删除时间
}
