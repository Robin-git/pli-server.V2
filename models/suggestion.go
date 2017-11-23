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
	EtablishmentID uint          `gorm:"not null" json:"etablishment_id"`
	UserID         uint          `gorm:"not null" json:"user_id"`
	ItemID         uint          `gorm:"null" json:"item_id"`
	Etablishment   *Etablishment `gorm:"ForeignKey:etablishment_id" json:"etablishment"`
	User           *User         `gorm:"ForeignKey:user_id" json:"user"`
	Item           *Item         `gorm:"ForeignKey:item_id" json:"item"`
}

// Suggestions is list of Suggestion
type Suggestions []*Suggestion

// QueryParameterSuggestion is list of possible parameter
type QueryParameterSuggestion struct {
	Etablishment bool
	User         bool
	Item         bool
}

func buildQueryWithEtablishment(db *gorm.DB, qp *QueryParameterSuggestion) *gorm.DB {
	if qp.Etablishment {
		return db.Preload("Etablishment")
	}
	return db
}

func buildQueryWithUser(db *gorm.DB, qp *QueryParameterSuggestion) *gorm.DB {
	if qp.User {
		return db.Preload("User")
	}
	return db
}

func buildQueryWithItem(db *gorm.DB, qp *QueryParameterSuggestion) *gorm.DB {
	if qp.Item {
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
func (s *ServiceSuggestion) GetSuggestions(qp *QueryParameterSuggestion, etablishmentID uint) (*Suggestions, error) {
	sg := &Suggestions{}
	db := buildQueryWithParam(s.DB, qp)
	return sg, db.Where(&Suggestion{EtablishmentID: etablishmentID}).Find(sg).Error
}

// GetSuggestion get one suggestion by id
func (s *ServiceSuggestion) GetSuggestion(id int, qp *QueryParameterSuggestion) (*Suggestion, error) {
	sg := &Suggestion{}
	db := buildQueryWithParam(s.DB, qp)
	return sg, db.First(sg).Error
}

// PostSuggestion post one suggestion
func (s *ServiceSuggestion) PostSuggestion(body struct {
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description" binding:"required"`
	EtablishmentID uint   `json:"etablishment_id" binding:"required"`
	UserID         uint   `json:"user_id" binding:"required"`
}) error {
	item := &Item{}
	suggestion := &Suggestion{}

	item.Price = -1
	item.Name = body.Name
	item.Description = body.Description
	item.EtablishmentID = body.EtablishmentID

	suggestion.EtablishmentID = body.EtablishmentID
	suggestion.UserID = body.UserID

	err := s.DB.Create(item).Error
	if err != nil {
		return err
	}
	suggestion.ItemID = item.ID
	return s.DB.Create(suggestion).Error
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

// DeleteSuggestion delete one suggestion
func (s *ServiceSuggestion) DeleteSuggestion(id uint) error {
	return s.DB.Delete(Suggestion{}, "id = ?", id).Error
}
