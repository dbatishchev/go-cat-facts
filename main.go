package main

import (
	"net/http"
	"go-cat-facts/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.MainPageHandler)
	http.ListenAndServe(":8080", nil)
}

