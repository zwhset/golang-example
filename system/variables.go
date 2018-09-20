package main

import "fmt"

func main() {
	// var 声明一个或多个变量

	var a = "initial" // 类型推导 var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2 // 多赋值
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short" // 类型推导
	fmt.Println(f)
}
