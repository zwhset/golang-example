package main

import (
	"io"
	"log"
	"net"
	"os"
)

// <-done 接收语句，丢弃结果，也就是说done里面有数据的库就停止阻塞
// ch <- x 发送语句
// x = <-ch 接收语句并赋值给x
// Close(ch) 关闭通道，设定标志拒绝发送请求， 但可以获取通道的数据。
// ch = make(chan int, 0/3) 0 无缓冲default, 3容量为3的缓冲通道
// x, ok := <-naturals !ok 表示通过关闭并且已经读完
func main() {
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		// 把连接的数据 Copy到标准输出中，并没有处理标准输入
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Fatal(err)
		}
		log.Println("done")
		done <- struct{}{} // tell main goroutine
	}()

	// 把标准输入 Copy到连接当中
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait goroutine done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
