package models

import (
	"fmt"

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
func (s *Service) GetUsers() {
	var users []User
	s.DB.Find(&users)
	fmt.Println(users)
}
