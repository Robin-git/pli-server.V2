package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceItem is a service of Item
type ServiceItem struct {
	DB *gorm.DB
}

//Item (id, name, price, description)
type Item struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `gorm:"not null" binding:"required" json:"name"`
	Price          float64   `gorm:"null" binding:"required" json:"price"`
	Description    string    `gorm:"not null" binding:"required" json:"description"`
	EtablishmentID uint      `gorm:"index:etablishment_id" json:"etablishment_id"`
}

// Items is list of Item
type Items []*Item

// GetItems get all item
func (s *ServiceItem) GetItems(idEtablishment uint) (*Items, error) {
	items := &Items{}
	return items, s.DB.Where(&Item{EtablishmentID: idEtablishment}).Not("price = ?", -1).Find(items).Error
}

// GetItem get one item
func (s *ServiceItem) GetItem(id int) (*Item, error) {
	item := &Item{}
	return item, s.DB.First(item, id).Error
}

// PostItem post one item
func (s *ServiceItem) PostItem(item *Item) error {
	return s.DB.Create(item).Error
}

// UpdateItem update one item
func (s *ServiceItem) UpdateItem(item *Item, itemUpdated *Item) error {
	item.Name = itemUpdated.Name
	item.Price = itemUpdated.Price
	item.Description = itemUpdated.Description
	item.EtablishmentID = itemUpdated.EtablishmentID
	return s.DB.Save(item).Error
}

// DeleteItem delete one item
func (s *ServiceItem) DeleteItem(id uint) error {
	return s.DB.Delete(Item{}, "id = ?", id).Error
}
