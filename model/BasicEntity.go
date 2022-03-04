package model

import (
	"time"

	"gorm.io/gorm"
)

type BasicEntity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
