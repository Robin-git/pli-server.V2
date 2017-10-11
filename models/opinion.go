package models

import "github.com/jinzhu/gorm"

//Opinion (id, User.id, Etablishment.id, *note*, *comment*, date)
type Opinion struct {
	gorm.Model
	Comment        string
	Note           int
	EtablishmentID uint `gorm:"not null" binding:"required"`
	UserID         uint `gorm:"not null" binding:"required"`
}

type Opinions []Opinion

func (s *Service) GetOpinion(idEtablishment uint) (*Opinions, error) {
	opinions := &Opinions{}
	return opinions, s.DB.Where(&Opinion{EtablishmentID: idEtablishment}).Find(opinions).Error
}

func (s *Service) PostOpinion(opinion Opinion) error {
	return s.DB.Create(&opinion).Error
}
