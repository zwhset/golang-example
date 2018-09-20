package main

import "fmt"

func main() {
	a, b, c, x := 1, 2, 3, 2

	// 普通switch, 与C语言不同不需要break
	switch x {
	case a, b: // or 命中一个即可
		fmt.Println(a | b)
	case c:
		fmt.Println("c")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("default")
	}

	// 初始化switch
	switch x := 9; x {
	case 1:
		fmt.Println("1")
	case 3:
		fmt.Println("3")
	case 5:
		fmt.Println("5")
	case 9: // 无语句块即单条件 case 9: break;

	default:
		fmt.Println("Nil")
	}

	// fallthrough 执行后续
	switch x := 1; x {
	default:
		fmt.Println("default")
	case 0:
		fmt.Println("0")
	case 1: // 匹配执行下一分支 即直接执行 case 2中代码
		fmt.Println("fallthrougth 1")
		fallthrough
	case 2:
		fmt.Println("fallthrougth 2")
		fallthrough // 没有下一分支即会报错，按源码顺序执行
	case 3:
		fmt.Println("fallthrougth 3")
	}

	// 替换if
	switch x := 5; {
	case x > 5:
		fmt.Println("x > 5")
	case x == 5:
		fmt.Println("x == 5")
	case x < 5:
		fmt.Println("x < 5")
	default:
		fmt.Println("default")
	}

}
