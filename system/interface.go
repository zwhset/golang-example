package main

// 设计接口时不要一上来就创建一系列接口， 可以用导出机制来限制一个类型的哪些方法或结构体的哪些字段是对包外可凶的
// 公在有两个或多个具体类型需要按统一的方式处理时才需要接口
// 另外接口和类型实现出于依赖的原因不能放在同一个包里，那么一个接口只有一个具体类型实现也是可以的，用接口来解o

import (
	"fmt"
	"io"
	"os"
	"syscall"
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
type data struct{}

func (*data) test() {}

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

	// 类型断言
	var w io.Writer
	w = os.Stdout // w 赋值给 os.Stdout类型

	f := w.(*os.File) // 成功： f == os.Stdout 即 os
	//c := w.(*bytes.Buffer) // 错误 w接口持有的是*os.File 不是 *bytes.Buffer

	rw := w.(io.ReadWriter) // 成功 *os.File 有Read和Write方法

	fmt.Println(f, rw)

	// 类型判断 判断w接口的类型是哪个
	if _, ok := w.(*os.File); ok {
		// do something
		fmt.Println(ok)
	}

	// 下面三个错误是写程序需要处理的, 可以使用到pathError
	//os.IsExist() 创建文件时文件已经存在
	//os.IsNotExist() 读取文件时不存在
	//os.IsPermission() 权限不足

	//type pathError struct {
	//	Op   string // 操作方式
	//	Path string // 路径
	//	Err  error  // 错误类型，里面包含 Err.Error()方法
	//}

	_, err := os.Open("/tmp/no/suchfile")
	fmt.Printf("%#v\n", err)
	// output &os.PathError{Op:"open", Path:"/tmp/no/suchfile", Err:0x2}

	// 如果pathError.Err == os.ErrNotExist 即不存在
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	fmt.Printf("File NotExist: %t", err == syscall.ENOENT || err == os.ErrNotExist)

	// 类型分支
	type X interface{}

}

func typeBranch(x interface{}) string {
	switch x.(type) {
	case nil:
	case int, uint:
	case bool:
	case string:
	default:
	}
}
