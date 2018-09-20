package main

import (
	"fmt"
)

// 接口(interface)是命名了的方法签名的集合
// 接口代表一种调用契约，是多个方法声明的集合，其它动态语言中也称协议。
// 准备交互的双方，共同遵守先约定的规则，使用在无须知道对方身份的情况下进行协作
// 好处 解除了类型依赖，有助于减少用户可视方法、屏蔽内部结构和实现细节

//type iface interface {
//	tab *itab
//	data unsafe.Pointer
//}

// 定义了一个接口
type tester interface {
	test()
	string() string
}

// data结构体类型实现了tester接口，因为实现了tester接口的方法
type data struct {}

func (*data) test()  {}

func (data) string() string {
	return ""
}


func main() {
	// 接口
	var d data
	var t tester = &d
	t.test()
	println(t.string())

	// 接口相等性
	fmt.Println("Interface type equal")
	var t1, t2 interface{}
	fmt.Println(t1 == nil, t1 == t2)

}