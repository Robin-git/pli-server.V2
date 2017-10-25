package models

import (
	"github.com/jinzhu/gorm"
)

// ServiceEtablishment is a service of Etablishment
type ServiceEtablishment struct {
	*Database
}

// Etablishment structure
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

// Etablishments is list of Etablishment
type Etablishments []Etablishment

// EtablishmentExtended is Etablishment whit
// - NoteAverage
type EtablishmentExtended struct {
	Etablishment
	NoteAverage float64
}

// EtablishmentExtendeds is list of EtablishmentExtended
type EtablishmentExtendeds []EtablishmentExtended

// GetEtablishments return all etablishments
func (s *ServiceEtablishment) GetEtablishments() (interface{}, error) {
	result := &EtablishmentExtendeds{}
	etablishments := &Etablishments{}
	if err := s.DB.Preload("Opinions").Find(&etablishments).Error; err != nil {
		return result, err
	}
	for _, e := range *etablishments {
		id := int(e.ID)
		noteAverage, _ := s.GetAverageNoteEtablishment(id)
		r := &EtablishmentExtended{
			e,
			noteAverage.Note,
		}
		*result = append(*result, *r)
	}
	return result, nil
}

// GetEtablishment return one etablishment
func (s *ServiceEtablishment) GetEtablishment(id int) (*Etablishment, error) {
	etablishment := &Etablishment{}
	return etablishment, s.DB.Preload("Opinions").First(etablishment, id).Error
}

// GetDistanceEtablishment return distance from position user and etablishment x and y
func (s *ServiceEtablishment) GetDistanceEtablishment(x, y, dist float64) (*Etablishments, error) {
	etablishments := &Etablishments{}
	var (
		query = `select * 
		from (
			select *, ((sqrt((pow((x - ?),2))+(pow((y - ?),2)))*1000) / 25) as distance 
			from etablishment 
		) as result where distance < ? order by distance`
	)
	return etablishments, s.DB.Raw(query, x, y, dist).Scan(etablishments).Error
}

// SearchEtablishmentByName search etablishment by name
func (s *ServiceEtablishment) SearchEtablishmentByName(r string) (*Etablishments, error) {
	etablishments := &Etablishments{}
	var (
		query = `select * from etablishment where replace(name, "-", " ") LIKE ?`
	)
	return etablishments, s.DB.Raw(query, r+"%").Scan(&etablishments).Error
}

// GetAverageNoteEtablishment return a average note from one etablishment
func (s *ServiceEtablishment) GetAverageNoteEtablishment(id int) (struct{ Note float64 }, error) {
	var (
		average struct {
			Note float64
		}
		query = `select avg(note) as note 
				from etablishment, opinion 
				where etablishment.id = ? 
				and etablishment.id = opinion.etablishment_id`
	)
	return average, s.DB.Raw(query, id).Scan(&average).Error
}
