package main

import "fmt"

func main() {
	n := 5

	// 条件循环
	for n <= 10 {
		fmt.Println(n)
		n++
	}

	// 死循环
	//for {
	//	fmt.Println("Loop")
	//}

	// 初始化表达式
	// 初始化值， 条件， 循环后
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	// break continue
	for j := 0 ;j <= 5; j++ {
		if j == 3 { // j == 3 时跳过此循环执行下一轮
			continue
		}
		fmt.Println(j)
	}
	// ouput 0 1 2 4 5

	for z := 0; z<=5; z++ {
		if z == 3 { // j == 3 时跳出循环
			break
		}
		fmt.Println(z)
	}
	// output 0 1 2

}
