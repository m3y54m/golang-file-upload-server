package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
)

func createImage(w http.ResponseWriter, r *http.Request) {

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(requestDump))

	fmt.Println("File Upload Endpoint Hit")

	if r.Method == "POST" {
		// Parse our multipart form, 10 << 20 ( results 10 * 1024 * 1024) specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handler, err := r.FormFile("myImage")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Create a temporary file within our directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("./", "file-*.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

	} else {
		fmt.Fprintf(w, "Invalid Request Method. Use POST instead.")
		return
	}

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func main() {
	createImageHandler := http.HandlerFunc(createImage)
	http.Handle("/", createImageHandler)
	port := os.Getenv("PORT")
	fmt.Println("Starting server on port " + port)
	http.ListenAndServe(":"+port, nil)
}
