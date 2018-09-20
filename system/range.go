package main

import "fmt"

func main() {

	// range slice or list
	l1 := [5]int{1, 2, 3, 4, 5}
	s1 := []string{"a", "b", "c", "d", "e"}
	//m1 := make(map[string]int)
	m1 := map[string]int{
		"user": 1,
		"pass": 2,
	}

	// for list
	for i, v := range l1 {
		fmt.Println("list: ", i, v)
	}

	// for slice
	for i, v := range s1 {
		fmt.Println("Slice: ", i, v)
	}

	// for map
	for k, v := range m1 {
		fmt.Println("Map: ", k, v)
	}
}
