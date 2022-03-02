package model

import (
	"time"

	"gorm.io/gorm"
)

type BasicEntity struct {
	ID        uint           `gorm:"primarykey;autoIncrement;comment:'主键id'"`
	CreatedAt time.Time      `gorm:"not null;autoCreateTime:nano;comment:'创建时间';"`                // 创建时间 默认时间为当前时间
	UpdatedAt time.Time      `gorm:"not null;autoCreateTime:nano;autoUpdateTime;comment:'更新时间';"` // 更新时间 默认时间为当前时间 当更新时，更新时间也更新
	DeletedAt gorm.DeletedAt `gorm:"index;comment:'删除时间';"`                                       // 删除时间
}
