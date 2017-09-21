package api

import (
	"gloo-server/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter return new mux Router
func NewRouter() *gin.Engine {
	// MariaDB is database of gloo app
	MariaDB := &SQLConnection{
		Login:    os.Getenv("GLOO_LOGIN"),
		Password: os.Getenv("GLOO_PWD"),
		Database: os.Getenv("GLOO_DATABASE"),
	}
	service := MariaDB.GetConnection()
	mariaController := controllers.ControllerScoped(service)

	r := gin.Default()

	r.GET("/api/user", mariaController.HandlerGetUsers)
	r.GET("/api/user/:id", mariaController.HandlerGetUser)

	r.GET("/api/etablishment", mariaController.HandlerGetEtablishments)
	r.GET("/api/etablishment/:id", mariaController.HandlerGetEtablishment)
	r.GET("/api/distance", mariaController.HandlerGetDistanceEtablishment)

	r.GET("/api/opinion", mariaController.HandlerGetOpinions)
	r.POST("/api/opinion", mariaController.HandlerPostOpinion)

	return r
}
