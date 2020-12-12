package main

import (
	"ch12/search"
	"net/http"
)

func main()  {
	http.HandleFunc("/search", search.Search)
	http.ListenAndServe(":12345", nil)
}
