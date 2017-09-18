package main

import (
	"fmt"
	"net/http"

	"gloo/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	mariadb := models.InitMariaDB()

	var user []models.User
	mariadb.Find(&user)
	fmt.Println(user)

	// var newUser = models.User{FirstName: "Robin", LastName: "Guerin", Email: "machin@gmail.com"}
	// db.Create(&newUser)

	mariadb.Find(&user)
	fmt.Println(user)

	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// log.Fatal(http.ListenAndServe(":8000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
