package models

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	UserId string `gorm:"not null;index"`
	Title  string `gorm:"not null"`
}
