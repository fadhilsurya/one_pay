package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName    string `gorm:"not null"`
	Username    string `gorm:"unique;not null"`
	Address     string `gorm:"not null;default:''"`
	PhoneNumber string `gorm:"unique;not null"`
	Role        string `gorm:"not null;default:'user'"`
	UserCode    string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
}
