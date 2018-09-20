package main

// GO支持在结构体类型中定义方法(methods)实现面向对象的概念
// 方法是面向对象的基本概述，用于维护和展示对象的自身状态
// 可以为当前包，除了接口和指针以外的任何类型定义方法

import (
	"strings"
	"fmt"
)


// 定义一个类型
type Str string

// 类型的方法 (s Str)前置实例接收参数
func (s Str) ToUpper() string {
	return strings.ToUpper(string(s))
}

// 类型的方法
func (s Str) ToLowwer() string {
	return strings.ToLower(string(s))
}

func main() {
	var s Str = "Zwhset"
	fmt.Println(s.ToUpper())
	fmt.Println(s.ToLowwer())
}
