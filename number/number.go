package main

import (
	"strconv"
	"fmt"
)

func main() {
	// string to float
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// string to int
	i, _ := strconv.ParseInt("1.234", 0, 64)
	fmt.Println(i)

	// 16 -> int
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// string to uint
	u, _ := strconv.ParseUint("786", 0, 64)
	fmt.Println(u)

	// string to number
	k, _ := strconv.Atoi("798")
	fmt.Println(k)

}
