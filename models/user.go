package models

import (
	"github.com/jinzhu/gorm"
)

// ServiceUser is a service of User
type ServiceUser struct {
	DB *gorm.DB
}

// User is struct of users
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string    `gorm:"not null"`
	Role      Role      `gorm:"ForeignKey:user_id"`
	Opinions  []Opinion `gorm:"ForeignKey:user_id"`
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
func (s *ServiceUser) GetUser(id int) (User, error) {
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
