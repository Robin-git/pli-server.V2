package api

import (
	"fmt"
	"gloo-server/models"
	"log"
	"os"
)

// SQLConnection is type of connection
type SQLConnection struct {
	Host     string
	Login    string
	Password string
	Database string
	Port     string
}

// GetConnection return service whith connection to sql database
func (c *SQLConnection) GetConnection() *models.Service {
	var env string
	switch os.Getenv("GLOO_ENV") {
	case "dev":
		env = fmt.Sprintf("%s_dev", c.Database)
		log.Println("Started on development")
	case "rec":
		env = fmt.Sprintf("%s_rec", c.Database)
		log.Println("Started on recipe")
	case "prod":
		env = fmt.Sprintf("%s_prod", c.Database)
		log.Println("Started on production")
	default:
		env = fmt.Sprintf("%s_dev", c.Database)
		log.Println("Warning : No GLOO_ENV selected")
		log.Println("Default --> Started on development")
	}
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.Login, c.Password, c.Host, c.Port, env)
	return models.InitMariaDB(conn)
}
