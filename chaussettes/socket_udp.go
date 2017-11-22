package chaussettes

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

type etablishment struct {
	ID    int   `json:"id"`
	Value int64 `json:"value"`
}

// LauchUDPServer lauch a udp connection
func LauchUDPServer(r *redis.Client) {
	port := ":7778"
	udpAddr, e := net.ResolveUDPAddr("udp", port)
	checkError(e)
	listener, e := net.ListenUDP("udp", udpAddr)
	checkError(e)
	log.Println("Server socket udp started on port " + port)
	handlerUDP(listener, r) // Run listener socket
}

func handlerUDP(conn *net.UDPConn, r *redis.Client) {
	defer func() {
		conn.Close()
		log.Println("Server udp closed")
	}()

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)

		var bit []byte
		var result []*etablishment

		stringbuffer := strings.Replace(string(buf[0:n]), "\n", "", -1)
		fmt.Println("Received ", stringbuffer, " from ", addr)

		idl := strings.Split(stringbuffer, ",")
		for _, id := range idl {
			e := &etablishment{}
			e.ID, err = strconv.Atoi(id)
			if err != nil {
				log.Println("Error in convert id")
			}
			e.Value, err = r.SCard(fmt.Sprint("etablishment:" + id)).Result()
			if err != nil {
				log.Println("Error get data from redis")
			}
			result = append(result, e)
		}
		bit, _ = json.Marshal(result)
		conn.WriteToUDP(bit, addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
