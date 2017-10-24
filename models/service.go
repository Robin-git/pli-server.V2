package models

import "github.com/jinzhu/gorm"

// Database is struct of service
type Database struct {
	DB *gorm.DB
}

// Service contain all services
type Service struct {
	ServiceUser
	ServiceEtablishment
	ServiceOpinion
}
