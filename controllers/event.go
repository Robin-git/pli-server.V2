package controllers

import (
	"gloo-server/models"

	"github.com/jinzhu/gorm"
)

// CtrlScopedEvent is initialize new Controller
func CtrlScopedEvent(db *gorm.DB) *CtrlEvent {
	return &CtrlEvent{
		Service: &models.ServiceEvent{
			DB: db,
		},
	}
}

// CtrlEvent is controller of Event
type CtrlEvent struct {
	Service *models.ServiceEvent
}
