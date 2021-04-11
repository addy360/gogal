package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"not null;unique"`
	Pasword      string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}
