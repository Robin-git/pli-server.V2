package redis

import (
	"log"

	"github.com/go-redis/redis"
)

// NewClient init new client redis
func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "[plidebian.cloudapp.net]:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Connect to redis client fail")
	}
	log.Println("Connection to redis client good")
	return client
}
