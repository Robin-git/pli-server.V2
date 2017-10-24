package models

import (
	"github.com/jinzhu/gorm"
	// Dialect for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitMariaDB return connection to mariadb gloo
func InitMariaDB(c string) *Database {
	db, err := gorm.Open("mysql", c)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	// All Migration
	db.AutoMigrate(
		&User{},
		&Role{},
		&Etablishment{},
		&Opinion{},
	)

	return &Database{
		DB: db,
	}
}
