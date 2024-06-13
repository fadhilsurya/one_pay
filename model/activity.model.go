package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Request   string    `gorm:"not null;default:''"`
	UserCode  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}
