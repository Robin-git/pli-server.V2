package api

import (
	"fmt"
	"gloo-server/controllers"
	"gloo-server/models"
	"log"
	"os"

	"github.com/gorilla/mux"
)

// SQLConnection is type of connection
type SQLConnection struct {
	Login    string
	Password string
	Database string
}

// GetConnection return service whith connection to sql database
func (c *SQLConnection) GetConnection() *models.Service {
	var env string
	switch os.Getenv("GLOO_ENV") {
	case "dev":
		env = fmt.Sprintf("%s_dev", c.Database)
		log.Println("Started on dev")
	case "prod":
		env = fmt.Sprintf("%s_prod", c.Database)
		log.Println("Started on prod")
	default:
		env = fmt.Sprintf("%s_dev", c.Database)
		log.Println("Warning : No GLOO_ENV selected")
		log.Println("Default --> Started on dev")
	}
	conn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", c.Login, c.Password, env)
	return models.InitMariaDB(conn)
}

// NewRouter return new mux Router
func NewRouter() *mux.Router {
	// MariaDB is database of gloo app
	MariaDB := &SQLConnection{
		Login:    "root",
		Password: "admin",
		Database: "gloo",
	}
	service := MariaDB.GetConnection()
	mariaController := controllers.ControllerScoped(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/users", mariaController.HandlerGetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", mariaController.HandlerGetUser).Methods("GET")
	return r
}
