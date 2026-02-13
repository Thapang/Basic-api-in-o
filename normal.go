package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello %s\n", name)
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
