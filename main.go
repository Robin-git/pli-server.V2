package main

import (
	"gloo-server/api"
	"log"
)

func main() {
	r := api.NewRouter()
	log.Println("Server starting")

	r.Run(":8000")
}
