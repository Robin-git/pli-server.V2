package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceOpinion is a service of Opinion
type ServiceOpinion struct {
	DB *gorm.DB
}

//Opinion (id, User.id, Etablishment.id, *note*, *comment*, date)
type Opinion struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Comment        string    `json:"comment"`
	Note           int       `json:"note"`
	EtablishmentID uint      `gorm:"not null" binding:"required" json:"etablishment_id"`
	UserID         uint      `gorm:"not null" binding:"required" json:"user_id"`
}

// Opinions is list of Opinion
type Opinions []Opinion

// GetOpinion return one opinion
func (s *ServiceOpinion) GetOpinion(idEtablishment uint) (*Opinions, error) {
	opinions := &Opinions{}
	return opinions, s.DB.Where(&Opinion{EtablishmentID: idEtablishment}).Find(opinions).Error
}

// PostOpinion post one opinion
func (s *ServiceOpinion) PostOpinion(opinion Opinion) error {
	return s.DB.Create(&opinion).Error
}
