package main

import (
	"unsafe"
	"fmt"
)

// 切片本身并非动态数组或数组指针，它内部通过指针引用底层数组，
// 设定相关的属性将数据读写操作限定在指定区域内。
type slice struct {
	array unsafe.Pointer
	len int
	cap int
}

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	// change
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	// append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	// slice
	fmt.Println(s[2:5])
	fmt.Println(s[2:])
	fmt.Println("len: ", len(s))
	fmt.Println(s[2:len(s)])

	// 简化赋值
	t := []string{"g", "h", "i"}
	fmt.Println(t)


}
