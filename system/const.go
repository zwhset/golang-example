package main

import (
	"fmt"
	"math"
)

// 定义常量
const s string  = "CONST"

func main() {

	fmt.Println(s)

	const n = 5000 // 类型推导

	const d  = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d)) // 类型转换

	fmt.Println(math.Sin(n), math.Pi) // math module
}
