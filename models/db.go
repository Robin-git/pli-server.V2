package models

import (
	"github.com/jinzhu/gorm"
	// Dialect for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitMariaDB return connection to mariadb gloo
func InitMariaDB(c string) *Service {
	db, err := gorm.Open("mysql", c)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	// All Migration
	db.AutoMigrate(User{})
	db.AutoMigrate(Role{})
	db.AutoMigrate(Etablishment{})

	return &Service{
		DB: db,
	}
}
