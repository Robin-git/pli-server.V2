package models

import (
	"github.com/jinzhu/gorm"
)

// User is struct of users
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}

// GetUsers find all User
func (s *Service) GetUsers() []User {
	var users []User
	s.DB.Find(&users)
	return users
}

// GetUser find one User
func (s *Service) GetUser(id int) User {
	var user User
	s.DB.First(&user, id)
	return user
}
