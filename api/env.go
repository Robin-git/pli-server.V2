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

func init() {
	ok := true
	for ok {
		switch {
		case os.Getenv("DB_HOST") == "":
			log.Println("DB_HOST par défault")
			os.Setenv("DB_HOST", "localhost")
		case os.Getenv("DB_LOGIN") == "":
			log.Println("DB_LOGIN par défault")
			os.Setenv("DB_LOGIN", "root")
		case os.Getenv("DB_DATABASE") == "":
			log.Println("DB_DATABASE par défault")
			os.Setenv("DB_DATABASE", "gloo")
		case os.Getenv("DB_PORT") == "":
			log.Println("DB_PORT par défault")
			os.Setenv("DB_PORT", "3306")
		default:
			ok = false
		}
	}
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
