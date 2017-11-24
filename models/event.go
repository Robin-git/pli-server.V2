package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceEvent is a service of Event
type ServiceEvent struct {
	DB *gorm.DB
}

//Event (id, Etablishment.id, nom, description, date)
type Event struct {
	ID             uint          `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	Name           string        `gorm:"not null" binding:"required" json:"name"`
	Description    string        `gorm:"not null;type:text" binding:"required" json:"description"`
	Date           time.Time     `gorm:"not null" binding:"required" json:"date"`
	EtablishmentID uint          `gorm:"index:etablishment_id" json:"etablishment_id"`
	Etablishment   *Etablishment `json:"etablishment"`
}

// Events is list of Event
type Events []*Event

// QueryParameterEvent is list of possible parameter
type QueryParameterEvent struct {
	Wetablishment bool
}

// GetEvents get all event
func (s *ServiceEvent) GetEvents(q *QueryParameterEvent, idEtablishment uint) (*Events, error) {
	events := &Events{}
	if q.Wetablishment {
		return events, s.DB.Where(&Event{EtablishmentID: idEtablishment}).Preload("Etablishment").Find(events).Error
	}
	return events, s.DB.Where(&Event{EtablishmentID: idEtablishment}).Find(events).Error
}

// GetEvent get one event by id
func (s *ServiceEvent) GetEvent(id int, q *QueryParameterEvent) (*Event, error) {
	event := &Event{}
	if q.Wetablishment {
		return event, s.DB.Preload("Etablishment").First(event, id).Error
	}
	return event, s.DB.First(event, id).Error
}

// PostEvent post one event
func (s *ServiceEvent) PostEvent(event *Event) error {
	return s.DB.Create(event).Error
}

// UpdateEvent update one event
func (s *ServiceEvent) UpdateEvent(event *Event, eventUpdated *Event) error {
	event.Name = eventUpdated.Name
	event.Description = eventUpdated.Description
	event.Date = eventUpdated.Date
	event.EtablishmentID = eventUpdated.EtablishmentID
	return s.DB.Save(event).Error
}
