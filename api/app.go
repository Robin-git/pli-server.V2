package api

import (
	"gloo-server/controllers"
	"gloo-server/models"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	service := models.InitMariaDB("root@/gloo_dev?charset=utf8&parseTime=True&loc=Local")
	// defer service.DB.Close()

	controller := controllers.ControllerScoped(service)

	r := mux.NewRouter()
	r.HandleFunc("/api/users", controller.HandlerGetUsers)
	r.HandleFunc("/api/users/{id}", controller.HandlerGetUser)
	return r
}
