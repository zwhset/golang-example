package main

import (
	"fmt"
	"time"
)

// 当 main goroutine结束时所有的 goroutine都结束，所以spinner也相应结束
//
func main() {
	go spinner(1000 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	// OUTPUT
	// # time go run goroutine/spinner.go
	// Fibonacci(45) = 1134903170
	// go run goroutine/spinner.go  8.10s user 0.16s system 101% cpu 8.109 total
}

// loading...
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

// Fib
func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
