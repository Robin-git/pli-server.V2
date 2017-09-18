package models

import "github.com/jinzhu/gorm"

type Etablishment struct {
	gorm.Model
	Name         string  `gorm:"not null"`
	X            float64 `gorm:"not null"`
	Y            float64 `gorm:"not null"`
	PhoneNumber  string  `gorm:"type:varchar(15)"`
	Email        string  `gorm:"type:varchar(256)"`
	PostalCode   string  `gorm:"type:varchar(10); not null"`
	City         string  `gorm:"type:varchar(256); not null"`
	StreetNumber string  `gorm:"type:varchar(4); not null"`
	Street       string  `gorm:"type:varchar(256); not null"`
}
