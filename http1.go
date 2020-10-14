package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	for item, price := range db {
		_, err := fmt.Fprintf(w, "%s: %s\n", item, price)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	db := database{"shoes": 50, "lists": 5}
	err := http.ListenAndServe("localhost:8080", db)
	if err != nil {
		log.Fatal(err)
	}
}
