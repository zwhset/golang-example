package main

import (
	"fmt"
	"time"
)

func main() {
	n := make(chan struct{}, 1)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Printf("Start Job %d\n", i)
			time.Sleep(3 * time.Second)
			fmt.Printf("Success Job %d\n", i)
			n <- struct{}{}
		}(i)
	}

	// 不阻塞的结果是 main先结束
	for i := 0; i < 100; i++ {
		<-n
	}
}
