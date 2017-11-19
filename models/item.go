package models

import (
	"github.com/jinzhu/gorm"
)

// ServiceItem is a service of Item
type ServiceItem struct {
	DB *gorm.DB
}

//Item (id, name, price, description)
type Item struct {
	gorm.Model
	Name        string  `gorm:"not null" binding:"required" json:"name"`
	Price       float64 `gorm:"not null" binding:"required" json:"price"`
	Description string  `gorm:"not null" binding:"required" json:"description"`
}

// Items is list of Item
type Items []*Item

// GetItems get all item
func (s *ServiceItem) GetItems() (*Items, error) {
	items := &Items{}
	return items, s.DB.Find(items).Error
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
	return s.DB.Save(item).Error
}
