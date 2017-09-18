package models

import (
	"github.com/jinzhu/gorm"
)

// User is struct of users
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"not null"`
	Role      Role   `gorm:"ForeignKey:RoleID"`
	RoleID    uint
}

// GetUsers find all User
func (s *Service) GetUsers() ([]User, error) {
	var users []User
	err := s.DB.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

// GetUser find one User
func (s *Service) GetUser(id int) (User, error) {
	var user User
	err := s.DB.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// GetUser find one User
func (s *Service) AddUser(u *User) error {
	err := s.DB.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}
