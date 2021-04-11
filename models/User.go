package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `gorm:"not null;unique"`
	Pasword       string `gorm:"-"`
	PasswordHash  string `gorm:"not null"`
	Remember      string `gorm:"-"`
	RememberToken string `gorm:"not null;unique"`
}
