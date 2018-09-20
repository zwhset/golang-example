package main

import (
	"os"
	"fmt"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	args := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(args)
	// OUTPUT
	//# go run src/line_args/args.go a b c d e
	//[/var/folders/k3/l1rymj0x6gqg9fk5b6h9f5l00000gn/T/go-build458312512/b001/exe/args a b c d e]
	//[a b c d e]
	//c

}
