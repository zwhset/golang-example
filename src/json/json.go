package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个response1结构体，用于将json转成此结构体
type response1 struct {
	Page int
	Fruits []string
}

// 定义一个json结构体，用于将结构体转成json
type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// 布尔转换
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	// 整数转换
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	// 浮点数转换
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	// 字符串转换
	strB, _ := json.Marshal("zwhset")
	fmt.Println(string(strB))

	// slice转换
	slcD := []string{"apple", "peach", "pear"}
	sliB, _ := json.Marshal(slcD)
	fmt.Println(string(sliB))

	// 字典/HASH转换
	mapD := map[string]int{"apple": 1, "peach": 2}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// 结构化转JSON
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear",}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 结构化转JSON 注意key的大小写
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear",}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// JSON 转结构体\
	fmt.Println()
	res2 := &response2{} // 初始化结构体
	json.Unmarshal([]byte(res2B), res2) // 将json解释成结构体
	fmt.Println(res2)

	// OUTPUT
	//true
	//1
	//2.34
	//"zwhset"
	//["apple","peach","pear"]
	//{"apple":1,"peach":2}
	//{"Page":1,"Fruits":["apple","peach","pear"]}
	//{"page":1,"fruits":["apple","peach","pear"]}
	//
	//&{1 [apple peach pear]}
}
