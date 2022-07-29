package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Simple Go File Upload Server")
	setupRoutes()
}
