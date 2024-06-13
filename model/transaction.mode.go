package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserCode        string  `gorm:"not null"`
	TransactionCode string  `gorm:"unique;not null"`
	Amount          float64 `gorm:"not null;default:0"`
	PaymentMethod   string  `gorm:"not null"`
	PaymentStatus   string  `gorm:"not null;default:'pending'"`
	Currency        string  `gorm:"not null"`
}
