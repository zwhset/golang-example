package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 0 -> 100
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	// float 32 or 64 0.0 <= f < 1.0
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Float64())

}
