package main

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
)

// 返回错误
func eadd(x, y int) (n int, err error) {
	return x + y, errors.New("This is test error.")
}

// 自定义错误
type MyError struct {
	x, y int
}

func (MyError) Error() string {
	return "This is MyError."
}

func tMyError(x, y int) (int, error) {
	return 0, MyError{x, y}
}

func main() {
	// 一般函数调用方法
	v, err := eadd(1, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v, err)

	// 调用自定义错误
	_, err = tMyError(1, 100)
	fmt.Println(err, "err:-")
	if err != nil {
		switch e := err.(type) {
		case MyError:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println("Not Fund Error.")
		}
	}

}
