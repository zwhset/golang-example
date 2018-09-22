package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例如 连接中止等错误
			continue
		}
		handlerConn(conn) // 一次处理一个连接
	}
}

func handlerConn(c net.Conn) {
	defer c.Close()

	for {
		// 一秒写一次时间
		if _, err := io.WriteString(
			c, time.Now().Format("15:04:05\n")); err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
