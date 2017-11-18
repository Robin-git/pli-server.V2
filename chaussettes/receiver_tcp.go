package chaussettes

import (
	"log"
	"net"
	"strings"

	"github.com/go-redis/redis"
)

// LauchTCPServerReceiver lauch a tcp connection
func LauchTCPServerReceiver(r *redis.Client) {
	port := ":7777"
	tcpAddr, e := net.ResolveTCPAddr("tcp4", port)
	checkError(e)
	listener, e := net.ListenTCP("tcp", tcpAddr)
	checkError(e)
	log.Println("Sender server socket tcp started on port " + port)
	receiverHandler(listener, r) // Run listener socket
}

func receiverHandler(l *net.TCPListener, r *redis.Client) {
	defer func() {
		l.Close()
		log.Println("Sender server tcp closed")
	}()
	for {
		conn, _ := l.Accept()
		// Make a buffer to hold incoming data.
		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		b, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			continue
		}
		// Convert byte to string
		msg := string(buf[:b])
		// Exemple : 101;1
		log.Println("Receive message : " + msg)
		// push data to Redis base
		pushToEtablishment(msg, r)
		// Close the connection when you're done with it.
		conn.Close()
	}
}

func pushToEtablishment(msg string, r *redis.Client) {
	// req[0] contain etablishment id
	// req[1] contain user id
	req := strings.Split(msg, reqParser)
	// exemple: etablishment:101 : { 1, 2 }
	r.SAdd("etablishment:"+req[0], req[1]).Result()
}
