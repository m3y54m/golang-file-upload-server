package main

import (
	"fmt"
	"net/http"
)

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "File Uploader: %q", r.URL.Path)
}

func setupRoutes() {
	http.HandleFunc("/", fileUploadHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Simple Go File Upload Server")
	setupRoutes()
}
