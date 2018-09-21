package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item] // ok-boim
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "No Such item %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WelCome HTTP4 Server.")
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func main() {

	db := database{"shoes": 50, "socks": 5}

	// HandlerFunc -> HandleFunc 专门实现一个方法来做mux干的事情
	http.HandleFunc("/", db.index)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	//mux := http.NewServeMux()
	//mux.Handle("/list", http.HandlerFunc(db.list))
	//mux.Handle("/price", http.HandlerFunc(db.price))

	log.Fatal(http.ListenAndServe("localhost:8004", nil))
}
