package chaussettes

import (
	"log"
	"net"
	"strconv"

	"github.com/go-redis/redis"
)

// LauchTCPServer lauch a tcp connection
func LauchTCPServer(r *redis.Client) {
	port := ":7778"
	tcpAddr, e := net.ResolveTCPAddr("tcp4", port)
	checkError(e)
	listener, e := net.ListenTCP("tcp", tcpAddr)
	checkError(e)
	log.Println("Server socket tcp started on port " + port)
	handler(listener, r) // Run listener socket
	log.Println("Server tcp closed")
}

func handler(l *net.TCPListener, r *redis.Client) {
	defer l.Close()
	counter := 0
	for {
		counter++
		conn, _ := l.Accept()
		log.Println(counter)
		conn.Write([]byte(strconv.Itoa(counter)))
		log.Println([]byte(strconv.Itoa(counter)))
		conn.Close()
	}

	// err := redis.Set("key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// val, err := redis.Get("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)
}
