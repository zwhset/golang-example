package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8001", db))
}

type database map[string]dollars

type dollars float32

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	for item, price := range db{
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}