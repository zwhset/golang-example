package main

import (
	"fmt"
	"time"
)

// 工作者 消费者
func worker(id int, jobs <- chan int, result chan <- int)  {
	for j := range jobs {
		fmt.Printf("worker[%d]: %d started job.\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker[%d]: %d finished job.\n", id, j)

		result <- j * 2
	}
}

func main() {

	// 定义一个任务列表与一个接收结果集
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 3个goroutine 工作，id分别是1，2，3
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 生产者
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 处理结果
	for a := 1; a <= 5; a++ {
		<- results
	}


}
