package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceSuggestion is a service of Suggestion
type ServiceSuggestion struct {
	DB *gorm.DB
}

// Suggestion (id, Etablishment.id, "User.id", Item.id)
type Suggestion struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Name           string        `gorm:"not null" binding:"required" json:"name"`
	// Price          int           `gorm:"not null" binding:"required" json:"price"`
	EtablishmentID int           `gorm:"not null" json:"etablishment_id"`
	UserID         int           `gorm:"not null" json:"user_id"`
	ItemID         int           `gorm:"not null" json:"item_id"`
	Etablishment   *Etablishment `gorm:"ForeignKey:etablishment_id" json:"etablishment"`
	User           *User         `gorm:"ForeignKey:user_id" json:"user"`
	Item           *Item         `gorm:"ForeignKey:item_id" json:"item"`
}

// Suggestions is list of Suggestion
type Suggestions []*Suggestion

// QueryParameterSuggestion is list of possible parameter
type QueryParameterSuggestion struct {
	WEtablishment bool
	WUser         bool
	WItem         bool
}

func buildQueryWithEtablishment(db *gorm.DB, qp *QueryParameterSuggestion) *gorm.DB {
	if qp.WEtablishment {
		return db.Preload("Etablishment")
	}
	return db
}

func buildQueryWithUser(db *gorm.DB, qp *QueryParameterSuggestion) *gorm.DB {
	if qp.WUser {
		return db.Preload("User")
	}
	return db
}

func buildQueryWithItem(db *gorm.DB, qp *QueryParameterSuggestion) *gorm.DB {
	if qp.WItem {
		return db.Preload("Item")
	}
	return db
}

func buildQueryWithParam(db *gorm.DB, qp *QueryParameterSuggestion) *gorm.DB {
	db1 := buildQueryWithEtablishment(db, qp)
	db2 := buildQueryWithUser(db1, qp)
	db3 := buildQueryWithItem(db2, qp)
	return db3
}

// GetSuggestions get all event
func (s *ServiceSuggestion) GetSuggestions(qp *QueryParameterSuggestion) (*Suggestions, error) {
	sg := &Suggestions{}
	db := buildQueryWithParam(s.DB, qp)
	return sg, db.Find(sg).Error
}

// GetSuggestion get one suggestion by id
func (s *ServiceSuggestion) GetSuggestion(id int, qp *QueryParameterSuggestion) (*Suggestion, error) {
	sg := &Suggestion{}
	db := buildQueryWithParam(s.DB, qp)
	return sg, db.First(sg).Error
}

// PostSuggestion post one suggestion
func (s *ServiceSuggestion) PostSuggestion(sg *Suggestion) error {
	return s.DB.Create(sg).Error
}

// UpdateSuggestion update one suggestion
func (s *ServiceSuggestion) UpdateSuggestion(sg *Suggestion, sgUpdated *Suggestion) error {
	// sg.Name = sgUpdated.Name
	// sg.Price = sgUpdated.Price
	sg.Etablishment = sgUpdated.Etablishment
	sg.User = sgUpdated.User
	sg.Item = sgUpdated.Item
	return s.DB.Save(sg).Error
}
