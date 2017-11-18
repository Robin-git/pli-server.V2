package models

import "github.com/jinzhu/gorm"

// ServiceSuggestion is a service of Suggestion
type ServiceSuggestion struct {
	DB *gorm.DB
}

// Suggestion (id, Etablishment.id, "User.id", Item.id)
type Suggestion struct {
	gorm.Model
	Name         string        `gorm:"not null" binding:"required"`
	Price        int           `gorm:"not null" binding:"required"`
	Etablishment *Etablishment `gorm:"ForeignKey:etablishment_id"`
	User         *User         `gorm:"ForeignKey:user_id"`
	Item         *Item         `gorm:"ForeignKey:item_id"`
}

// Suggestions is list of Suggestion
type Suggestions []*Suggestion
