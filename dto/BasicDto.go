package dto

import (
	"time"
)

type BasicDto struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}
