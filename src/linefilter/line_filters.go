package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stdout, "Error", err)
		os.Exit(1)
	}

	// OUTPUT
	//echo "zwhset" | go run src/linefilter/line_filters.go
	//ZWHSET

}
