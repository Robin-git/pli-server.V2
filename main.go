package main

import (
	"gloo-server/api"
	"gloo-server/chaussettes"
	"gloo-server/redis"
	"log"

	_ "net/http/pprof"
)

func main() {
	r := api.NewRouter()
	log.Println("API server starting")
	redis := redis.NewClient()

	go chaussettes.LauchTCPServerReceiver(redis)
	go chaussettes.LauchUDPServer(redis)

	// Server for analyse performance
	// go http.ListenAndServe(":8080", http.DefaultServeMux)
	r.Run(":8000")
}
