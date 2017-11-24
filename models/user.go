package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceUser is a service of User
type ServiceUser struct {
	DB *gorm.DB
}

// User is struct of users
type User struct {
	ID        string     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `gorm:"not null" json:"email"`
	Role      *Role      `gorm:"ForeignKey:user_id" json:"role"`
	Opinions  []*Opinion `gorm:"ForeignKey:user_id" json:"opinions"`
}

// GetUsers find all User
func (s *ServiceUser) GetUsers() ([]User, error) {
	var users []User
	err := s.DB.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

// GetUser find one User
func (s *ServiceUser) GetUser(id string) (User, error) {
	var user User
	err := s.DB.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// AddUser post one User
func (s *ServiceUser) AddUser(u *User) error {
	err := s.DB.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}
