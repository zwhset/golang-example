package main

import (
	"strings"
	"fmt"
)

// 实现类python 的字符串对象方法

type str string

func main() {
	var apple str = "apple"
	fmt.Println(apple.Len())
	fmt.Println(apple.ToUpper())
	fmt.Println(apple.ToLower())
}

func (s str) Len() int  {
	return len(s)
}

func (s str) ToUpper() string {
	return strings.ToUpper(string(s))
}

func (s str) ToLower() string {
	return strings.ToLower(string(s))
}

