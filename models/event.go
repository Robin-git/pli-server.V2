package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceEvent is a service of Event
type ServiceEvent struct {
	DB *gorm.DB
}

//Event (id, Etablishment.id, nom, description, date)
type Event struct {
	gorm.Model
	Name        string    `gorm:"not null" binding:"required"`
	Description string    `gorm:"not null" binding:"required"`
	Date        time.Time `gorm:"not null" binding:"required"`
}

// Events is list of Event
type Events []*Event
