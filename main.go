package main

import (
	"gloo-server/api"
	"log"
	"net/http"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := api.NewRouter()
	s := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server starting")
	log.Fatal(s.ListenAndServe())
}
