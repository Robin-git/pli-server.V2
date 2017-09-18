package models

import (
	"github.com/jinzhu/gorm"
)

// InitMariaDB return connection to mariadb gloo
func InitMariaDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:admin@/gloo_dev?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	return db
	// db.AutoMigrate(&models.User{})
	// defer db.Close()
}
