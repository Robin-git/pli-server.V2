package api

import (
	"gloo-server/controllers"

	"github.com/gorilla/mux"
)

// NewRouter return new mux Router
func NewRouter() *mux.Router {
	// MariaDB is database of gloo app
	MariaDB := &SQLConnection{
		Login:    "root",
		Password: "",
		Database: "gloo",
	}
	service := MariaDB.GetConnection()
	mariaController := controllers.ControllerScoped(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/users", mariaController.HandlerGetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", mariaController.HandlerGetUser).Methods("GET")
	return r
}
