package main

import (
	"gloo-server/controllers"
	"gloo-server/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	service := models.InitMariaDB()
	controller := controllers.ControllerScoped(service)

	// var newUser = models.User{FirstName: "Robin", LastName: "Guerin", Email: "machin@gmail.com"}
	// db.Create(&newUser)

	// mariadb.Find(&user)
	// fmt.Println(user)

	r := mux.NewRouter()
	r.HandleFunc("/api/users", controller.HandlerGetUsers)
	r.HandleFunc("/api/users/{id}", controller.HandlerGetUser)
	log.Fatal(http.ListenAndServe(":8000", r))
}
