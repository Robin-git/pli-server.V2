package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Value  string `gorm:"not null"`
	UserID uint
}
