package models

import (
	"github.com/jinzhu/gorm"
)

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
	if err := s.DB.Find(&etablishments).Error; err != nil {
		return []Etablishment{}, err
	}
	return etablishments, nil
}

func (s *Service) GetEtablishment(id int) (Etablishment, error) {
	var etablishment Etablishment
	if err := s.DB.First(&etablishment, id).Error; err != nil {
		return Etablishment{}, err
	}
	return etablishment, nil
}

func (s *Service) GetDistanceEtablishment(x, y, dist float64) ([]Etablishment, error) {
	var (
		r     []Etablishment
		query = `select * 
		from (
			select *, ((sqrt((pow((x - ?),2))+(pow((y - ?),2)))*1000) / 25) as distance 
			from gloo_rec.etablishment 
		) as result where distance < ? order by distance`
	)
	if err := s.DB.Raw(query, x, y, dist).Scan(&r).Error; err != nil {
		return []Etablishment{}, err
	}
	return r, nil
}

func (s *Service) SearchEtablishmentByName(r string) ([]Etablishment, error) {
	var (
		e     []Etablishment
		query = `select * from etablishment where name LIKE "%?%"`
	)
	if err := s.DB.Raw(query, r).Scan(&e).Error; err != nil {
		return []Etablishment{}, err
	}
	return e, nil
}
