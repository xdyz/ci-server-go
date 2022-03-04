package model

import (
	"time"

	"gorm.io/gorm"
)

type BasicEntity struct {
	Id        uint           `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间';"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:'更新时间';"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:'删除时间';"`
}
