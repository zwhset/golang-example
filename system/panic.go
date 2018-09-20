package main

import (
	"os"
	"fmt"
)

func main() {
	panic("a Problem")

	// 触发panic下面的代码将不会被执行
	_, err := os.Create("/tmp/panic.txt")
	if err != nil {
		panic(err)
	}

	// 触发panic defer语句也不会被执行
	defer fmt.Println("panic after defer not run.")
}