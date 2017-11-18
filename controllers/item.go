package controllers

import (
	"gloo-server/models"

	"github.com/jinzhu/gorm"
)

// CtrlScopedItem is initialize new Controller
func CtrlScopedItem(db *gorm.DB) *CtrlItem {
	return &CtrlItem{
		Service: &models.ServiceItem{
			DB: db,
		},
	}
}

// CtrlItem is controller of Item
type CtrlItem struct {
	Service *models.ServiceItem
}
