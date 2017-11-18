package chaussettes

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/go-redis/redis"
)

type etablishment struct {
	ID    string `json:"id"`
	Value int64  `json:"value"`
}

// LauchTCPServerSender lauch a tcp connection
func LauchTCPServerSender(r *redis.Client) {
	port := ":7778"
	tcpAddr, e := net.ResolveTCPAddr("tcp4", port)
	checkError(e)
	listener, e := net.ListenTCP("tcp", tcpAddr)
	checkError(e)
	log.Println("Receiver server socket tcp started on port " + port)
	senderHandler(listener, r) // Run listener socket
}

func senderHandleReq(conn net.Conn, r *redis.Client) {
	for {
		l, _ := r.Keys("*etablishment*").Result()
		var res []*etablishment
		for _, id := range l {
			v, _ := r.SCard(id).Result()

			res = append(res, &etablishment{
				ID:    id,
				Value: v,
			})
		}
		resjson, _ := json.Marshal(res)
		conn.Write(resjson)
		time.Sleep(2 * time.Second)
	}
}

func senderHandler(l *net.TCPListener, r *redis.Client) {
	defer func() {
		l.Close()
		log.Println("Server tcp closed")
	}()
	for {
		conn, _ := l.Accept()
		go senderHandleReq(conn, r)
	}
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
