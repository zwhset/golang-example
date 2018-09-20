package main

import (
	"fmt"
	"github.com/labstack/gommon/log"

	"errors"
)

func main() {
	x := 3

	// if elseif else
	if x > 5 {
		fmt.Println("a")
	} else if x < 5 && x > 0 {
		fmt.Println("b")
	} else {
		fmt.Println("z")
	}

	// 块局部变量或初始化函数
	y := 10

	if xinit(); y == 0 { // 初始化函数
		fmt.Println("a")
	}

	if a, b := 1, 2; a < b { // 定义局部变量
		fmt.Println("a < b")
	}

	// error 当中的用法
	if err := getERror(); err != nil {
		log.Fatal(err)
	}


}

func xinit()  {
	fmt.Println("xinit")
}

func getERror() error {
	return errors.New("Test error.")
}
