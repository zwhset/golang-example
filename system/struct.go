package main

import "fmt"

// 结构体是带类型的字段集合

// 定义一个结构体
type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20}) // 直接给值
	fmt.Println(person{name:"Alice", age:30}) // 按key给值
	fmt.Println(person{name:"zwhset"}) // 零值初始化

	fmt.Println(&person{name:"l9set", age:20}) // 指针

	s := person{name: "sean", age: 50}
	fmt.Println(s.name) // 访问字段

	sp := &s // 复制指针
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age, s.age) // 指针修改两处都生效
}
