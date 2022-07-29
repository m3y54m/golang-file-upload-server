package main

import (
	"fmt"
	"net/http"
)

func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\r\nHello World!, Here is %q\r\n", r.URL.Path)

	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Method is GET")
	case "POST":
		fmt.Fprintln(w, "Method is POST")
	default:
		fmt.Fprintln(w, "Sorry, only GET and POST methods are supported.")
	}
}

func setupServer() {
	http.HandleFunc("/", httpRequestHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Simple Go File Upload Server")
	setupServer()
}
