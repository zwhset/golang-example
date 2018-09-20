package main

import "fmt"

type user struct {
	Username string
	Password string
	Age int
}

func main() {
	var a [3]int // [3] 3 必须是非负数, 零值初始化
	var b [2]int
	// d1 = d2 不是同类型不能赋值

	fmt.Println(a, b)

	var c [4]int // 元素自动初始化为0
	d := [5]int{2, 5} // [2, 5, 0, 0, 0]
	e := [5]int{3: 1, 4 : 2} // 下标从0开始 [0, 0, 0, 1, 2]
	fmt.Println(c, d, e)

	// 结构体初始化
	f := [...]user{
		user{Username:"zwhset"},
		user{Username:"l0set"},
	}
	fmt.Printf("%#v\n", f)

	// 多维数组中只允许第一维度使用 "..."
	// len 与cap 都只返回第一维的信息
}
