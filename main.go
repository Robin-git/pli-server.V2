package main

import (
	"net/http"

	"gloo-server/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	mariadb := models.InitMariaDB()

	service := &models.Service{
		DB: mariadb,
	}

	service.GetUsers()

	// var newUser = models.User{FirstName: "Robin", LastName: "Guerin", Email: "machin@gmail.com"}
	// db.Create(&newUser)

	// mariadb.Find(&user)
	// fmt.Println(user)

	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// log.Fatal(http.ListenAndServe(":8000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
