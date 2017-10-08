package models

import (
	"github.com/jinzhu/gorm"
)

type Etablishment struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	X           float64 `gorm:"not null"`
	Y           float64 `gorm:"not null"`
	PhoneNumber string  `gorm:"type:varchar(15)"`
	Email       string  `gorm:"type:varchar(256)"`
	PostalCode  string  `gorm:"type:varchar(10); not null"`
	City        string  `gorm:"type:varchar(256); not null"`
	Street      string  `gorm:"type:varchar(256); not null"`
	Opinions    []Opinion
}

func (s *Service) GetEtablishments() (interface{}, error) {
	type EtablishmentWithAverage struct {
		Etablishment
		Note_average float64
	}
	type EtablishmentsWithAverageResult []EtablishmentWithAverage
	var (
		etablishments []Etablishment
		result        EtablishmentsWithAverageResult
	)
	if err := s.DB.Preload("Opinions").Find(&etablishments).Error; err != nil {
		return result, err
	}
	for _, etablishment := range etablishments {
		id := int(etablishment.ID)
		note_average, _ := s.GetAverageNoteEtablishment(id)
		r := EtablishmentWithAverage{
			etablishment,
			note_average.Note,
		}
		result = append(result, r)
	}
	return result, nil
}

func (s *Service) GetEtablishment(id int) (Etablishment, error) {
	var etablishment Etablishment
	if err := s.DB.Preload("Opinions").First(&etablishment, id).Error; err != nil {
		return Etablishment{}, err
	}
	return etablishment, nil
}

func (s *Service) GetDistanceEtablishment(x, y, dist float64) ([]Etablishment, error) {
	var (
		etablishments []Etablishment
		query         = `select * 
		from (
			select *, ((sqrt((pow((x - ?),2))+(pow((y - ?),2)))*1000) / 25) as distance 
			from etablishment 
		) as result where distance < ? order by distance`
	)
	if err := s.DB.Raw(query, x, y, dist).Scan(&etablishments).Error; err != nil {
		return []Etablishment{}, err
	}
	return etablishments, nil
}

func (s *Service) SearchEtablishmentByName(r string) ([]Etablishment, error) {
	var (
		etablishments []Etablishment
		query         = `select * from etablishment where replace(name, "-", " ") LIKE ?`
	)
	if err := s.DB.Raw(query, r+"%").Scan(&etablishments).Error; err != nil {
		return []Etablishment{}, err
	}
	return etablishments, nil
}

func (s *Service) GetAverageNoteEtablishment(id int) (struct{ Note float64 }, error) {
	var (
		average struct {
			Note float64
		}
		query = `select avg(note) as note 
				from etablishment, opinion 
				where etablishment.id = ? 
				and etablishment.id = opinion.etablishment_id`
	)
	if err := s.DB.Raw(query, id).Scan(&average).Error; err != nil {
		return average, err
	}
	return average, nil
}
