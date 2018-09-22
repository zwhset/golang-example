package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handlerConn1(conn)
	}
}

func handlerConn1(c net.Conn) {
	defer c.Close()

	for {
		if _, err := io.WriteString(
			c, time.Now().Format("15:04:05\n")); err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
