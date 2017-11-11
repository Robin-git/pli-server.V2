package models

import "github.com/jinzhu/gorm"

// ServiceOpinion is a service of Opinion
type ServiceOpinion struct {
	DB *gorm.DB
}

//Opinion (id, User.id, Etablishment.id, *note*, *comment*, date)
type Opinion struct {
	gorm.Model
	Comment        string
	Note           int
	EtablishmentID uint `gorm:"not null" binding:"required"`
	UserID         uint `gorm:"not null" binding:"required"`
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
