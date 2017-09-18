package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Value string `sql:"not null"`
}
