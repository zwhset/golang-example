package main

import (
	"flag"
	"fmt"
)

func main() {
	// 参数名 默认值 说明，最后初始化即可食用
	wordPtr := flag.String("word", "foo", "a String") // argsname, default_value, desc
	numPtr := flag.Int("numb", 38, "a Int")
	boolPtr := flag.Bool("fork", false, "a Bool")

	// 赋值的方式
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a String var")
	// 必要
	flag.Parse()

	fmt.Println("word: ", *wordPtr)
	fmt.Println("num: ", *numPtr)
	fmt.Println("bool: ", *boolPtr)
	fmt.Println("sval: ", svar)
	fmt.Println("tail: ", flag.Args())

	// OUTPUT
	//go run src/line_args/flag.go --word=china --numb=88 --fork=true -svar=zwhset a b
	//word:  china
	//num:  88
	//bool:  true
	//sval:  zwhset
	//tail:  [a b]
}