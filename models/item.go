package models

import "github.com/jinzhu/gorm"

// ServiceItem is a service of Item
type ServiceItem struct {
	DB *gorm.DB
}

//Item (id, name, price, description)
type Item struct {
	gorm.Model
	Name        string  `gorm:"not null" binding:"required"`
	Price       float64 `gorm:"not null" binding:"required"`
	Description string  `gorm:"not null" binding:"required"`
}

// Items is list of Item
type Items []*Item
