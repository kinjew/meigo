package common

import (
	"meigo/library/time"
)

// BaseModel 是数据库的基本结构体
type BaseModel struct {
	ID         int         `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                       // 主键
	CreatedAt  time.MTime  `gorm:"column:created_at;" json:"created_at" form:"created_at"`                          // 创建时间
	UpdatedAt  time.MTime  `gorm:"column:updated_at;" json:"updated_at" form:"updated_at"`                          // 更新时间
	DeletedAt  *time.MTime `gorm:"column:deleted_at;default:null;" json:"deleted_at" form:"deleted_at" sql:"index"` // 删除时间
	CreatedUID int64       `gorm:"column:created_uid;default:0;not null;" json:"created_uid" form:"created_uid"`    // 创建人
	UpdatedUID int64       `gorm:"column:updated_uid;default:0;not null;" json:"updated_uid" form:"updated_uid"`    // 更新人
}

// BaseModelV1 是数据库的基本结构体
type BaseModelV1 struct {
	ID        int `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"` // 主键
	CreatedAt int `gorm:"column:created_at;" json:"created_at" form:"created_at"`    // 创建时间
	UpdatedAt int `gorm:"column:updated_at;" json:"updated_at" form:"updated_at"`    // 更新时间
}

// BaseModelV2 是数据库的基本结构体
type BaseModelV2 struct {
	ID        int `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                       // 主键
	CreatedAt int `gorm:"column:created_at;" json:"created_at" form:"created_at"`                          // 创建时间
	UpdatedAt int `gorm:"column:updated_at;" json:"updated_at" form:"updated_at"`                          // 更新时间
	DeletedAt int `gorm:"column:deleted_at;default:null;" json:"deleted_at" form:"deleted_at" sql:"index"` // 删除时间
}
