package api

import (
	"gloo-server/controllers"

	"github.com/gin-gonic/gin"
)

// NewRouter return new mux Router
func NewRouter() *gin.Engine {
	// MariaDB is database of gloo app
	MariaDB := &SQLConnection{
		Login:    "root",
		Password: "",
		Database: "gloo",
	}
	service := MariaDB.GetConnection()
	mariaController := controllers.ControllerScoped(service)

	r := gin.Default()

	r.GET("/api/user", mariaController.HandlerGetUsers)
	r.GET("/api/user/:id", mariaController.HandlerGetUser)

	r.GET("/api/etablishment", mariaController.HandlerGetEtablishments)
	r.GET("/api/etablishment/:id", mariaController.HandlerGetEtablishment)

	return r
}
