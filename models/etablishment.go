package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceEtablishment is a service of Etablishment
type ServiceEtablishment struct {
	DB *gorm.DB
}

// Etablishment structure
type Etablishment struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Name        string     `gorm:"not null" json:"name"`
	X           float64    `gorm:"not null" json:"x"`
	Y           float64    `gorm:"not null" json:"y"`
	PhoneNumber string     `gorm:"type:varchar(15)" json:"phone_number"`
	Email       string     `gorm:"type:varchar(256)" json:"email"`
	PostalCode  string     `gorm:"type:varchar(10); not null" json:"postal_code"`
	City        string     `gorm:"type:varchar(256); not null" json:"city"`
	Street      string     `gorm:"type:varchar(256); not null" json:"street"`
	Opinions    []*Opinion `gorm:"ForeignKey:etablishment_id" json:"opinions"`
	Items       []*Item    `gorm:"ForeignKey:etablishment_id" json:"items"`
}

// Etablishments is list of Etablishment
type Etablishments []Etablishment

// EtablishmentExtended is Etablishment whit
// - NoteAverage
type EtablishmentExtended struct {
	Etablishment
	NoteAverage float64 `json:"note_average"`
}

// EtablishmentExtendeds is list of EtablishmentExtended
type EtablishmentExtendeds []EtablishmentExtended

// GetEtablishments return all etablishments
func (s *ServiceEtablishment) GetEtablishments() (interface{}, error) {
	result := &EtablishmentExtendeds{}
	etablishments := &Etablishments{}
	if err := s.DB.Preload("Opinions").Preload("Items").Find(&etablishments).Error; err != nil {
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
	return etablishment, s.DB.Preload("Opinions").Preload("Items").First(etablishment, id).Error
}

// GetDistanceEtablishment return distance from position user and etablishment x and y
func (s *ServiceEtablishment) GetDistanceEtablishment(x, y, dist float64) (*EtablishmentExtendeds, error) {
	etablishments := &Etablishments{}
	var (
		query = `select * 
		from (
			select *, ((sqrt((pow((x - ?),2))+(pow((y - ?),2)))*1000) / 25) as distance 
			from etablishment 
		) as result where distance < ? order by distance`
	)
	err := s.DB.Raw(query, x, y, dist).Scan(etablishments).Error
	if err != nil {
		return nil, err
	}
	result := &EtablishmentExtendeds{}
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
