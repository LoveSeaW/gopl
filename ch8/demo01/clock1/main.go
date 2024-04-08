package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("net.Dial failed, err = %v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("accepet failed, err = %v", err)
			continue
		}
		handler(conn)

	}

}

func handler(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Fatalf("io.WriteString failed, err := %v", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
