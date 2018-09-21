package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8002", db))
}

type database map[string]dollars

type dollars float32

// InterFace ServeHTTP(ResponseWriter, *Request)
func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item] // ok-boim
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "No Such item %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No Such Page: %s\n", r.URL)
	}
}

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}
