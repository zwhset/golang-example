package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	// set or add
	m["k1"] = 7
	m["k2"] = 8
	fmt.Println(m)

	// get
	fmt.Println(m["k2"], m["k"])

	// 赋值
	v1 := m["k1"]
	fmt.Println(v1)

	// delete
	delete(m, "k2")
	fmt.Println("map: ", m)

	// 存员关系，是否存在
	k, v := m["k1"]
	fmt.Println(k, v)

	// 简单赋值
	n := map[string]string{"name": "zwhset", "city": "china"}
	fmt.Println(n)

	// 匿名map
	m2 := map[int]struct{
		x int
	}{
		1 : {x: 100},
		2 : {x: 200},
	}
	fmt.Println(m2)

	// ok-idiom
	if _, ok := m["d"]; !ok {
		m["d"] = 9
		fmt.Println(m)
	}

	// for
	for k, v := range m {
		fmt.Printf("key[%s]= %d\n", k, v)
	}

	// 初始化
	//var m3 map[string]int // 未初始化
	//m3["k"] = 1 // 未初始化不能赋值
	var m4 = map[string]int{} // 初始化
	//m5 := make(map[string]int, 1000) // 初始化，预先创建空间
	m4["k"] = 2 // 可以赋值
}
