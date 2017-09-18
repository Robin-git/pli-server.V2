package main

import (
	"gloo-server/api"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
