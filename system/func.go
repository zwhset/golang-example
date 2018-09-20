package main

import (
	"fmt"
	"time"
	"log"
	"errors"
	"os"
)

// 参数与返回值 参数是值拷贝
func add(x, y int) int {
	return x + y
}

func add2(x, y *int) int {
	fmt.Printf("func: x: %p\ty: %p\n", &x, &y)
	return *x + *y
}

// 参数过多，重构成结构体
type serverOption struct {
	address string
	port int
	path string
	timeout time.Duration
	log *log.Logger
}

// 默认参数
func newOption() *serverOption {
	return &serverOption{
		address: "0.0.0.0",
		port: 8000,
		path: "/",
		timeout: 10,
		log: nil,
	}
}

// 变参
func test(s string, a ...int)  {
	fmt.Printf("%T %v\n", a, a)
}

// 返回值，如果有设置返回值必须要有return
func ret() int {
	fmt.Println("func ret")
	return 0
}

// 多返回值
func manyRet() (int, error)  {
	return 1, errors.New("GoLang error.")
}

// 隐式返回
func hiddenRet(x int) (z int, err int) {
	switch x {
	case 0:
		z = 0
		return
	case 1:
		z = 1
		err = 90
		return
	default:
		return
	}
}

// 闭包，返回一个函数
func test10(x int) func() {
	return func() {
		fmt.Println(x * 100)
	}
}

// 延迟调用， 直到当函数执行结束前才被执行
func twriteFile(filename string)  {
	_, err := os.Stat(filename)
	if err != nil {
		if _, err := os.Create(filename); err !=nil {
			log.Fatalf("Create File Fail, %s", err)
		}
	}

	fd, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Open File Error, %s", err)
	}

	fd.WriteString("test.")
	fd.Sync()

	// 延迟调用
	defer fd.Close()
}

func main() {
	fmt.Println(add(1, 2))
	x := 2
	y := 3
	fmt.Printf("x: %p\ty: %p\n", &x, &y)
	fmt.Println(add2(&x, &y))

	// struct 参数
	opt := newOption()
	fmt.Println(opt)

	// 变参
	test("zwhset", 1, 3, 9)

	// 返回值
	ret()
	if _, err := manyRet(); err != nil {
		//log.Fatal(err)
	}

	// 隐式返回
	x, y = hiddenRet(1)
	//fmt.Printf("func hiddenRet: %d\t%s", x, y)
	fmt.Println(x, y, "x")

	// 匿名函数
	go func(x, y int) int {
		return x + y
	}(1, 2)

	// 闭包
	t := test10(5)
	t() // ouput 500

	// 延迟调用 defer
	twriteFile("tmp.txt")
}
