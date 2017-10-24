package models

import "github.com/jinzhu/gorm"

// Role is struct of Role
type Role struct {
	gorm.Model
	Value  string `gorm:"not null"`
	UserID uint
}
