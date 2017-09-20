package models

import "github.com/jinzhu/gorm"

type Etablishment struct {
	gorm.Model
	Name         string    `gorm:"not null"`
	X            float64   `gorm:"not null"`
	Y            float64   `gorm:"not null"`
	PhoneNumber  string    `gorm:"type:varchar(15)"`
	Email        string    `gorm:"type:varchar(256)"`
	PostalCode   string    `gorm:"type:varchar(10); not null"`
	City         string    `gorm:"type:varchar(256); not null"`
	Street       string    `gorm:"type:varchar(256); not null"`
	Opinions     []Opinion `gorm:"ForeignKey:EtablishmentID;AssociationForeignKey:OpinionRefer"`
	OpinionRefer uint
}

func (s *Service) GetEtablishments() ([]Etablishment, error) {
	var etablishments []Etablishment
	err := s.DB.Find(&etablishments).Error
	if err != nil {
		return []Etablishment{}, err
	}
	return etablishments, nil
}

func (s *Service) GetEtablishment(id int) (Etablishment, error) {
	var etablishment Etablishment
	err := s.DB.First(&etablishment, id).Error
	if err != nil {
		return Etablishment{}, err
	}
	return etablishment, nil
}
