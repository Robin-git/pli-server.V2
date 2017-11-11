package chaussettes

import (
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis"
)

// LauchUDPServer lauch a udp connection
func LauchUDPServer(r *redis.Client) {
	port := ":7777"
	udpAddr, e := net.ResolveUDPAddr("udp", port)
	checkError(e)
	listener, e := net.ListenUDP("udp", udpAddr)
	checkError(e)
	log.Println("Server socket udp started on port " + port)
	handlerUDP(listener, r) // Run listener socket
	log.Println("Server udp closed")
}

func handlerUDP(conn *net.UDPConn, r *redis.Client) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		bit := []byte("coucou\n")
		conn.WriteToUDP(bit, addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
