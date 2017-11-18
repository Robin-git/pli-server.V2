package controllers

import (
	"gloo-server/models"

	"github.com/jinzhu/gorm"
)

// CtrlScopedSuggestion is initialize new Controller
func CtrlScopedSuggestion(db *gorm.DB) *CtrlSuggestion {
	return &CtrlSuggestion{
		Service: &models.ServiceSuggestion{
			DB: db,
		},
	}
}

// CtrlSuggestion is controller of Suggestion
type CtrlSuggestion struct {
	Service *models.ServiceSuggestion
}
