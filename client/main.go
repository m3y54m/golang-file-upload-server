package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	call("http://localhost:80/", "POST")
}

func call(urlPath, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("myImage", "uploaded_image.jpg")
	if err != nil {
		fmt.Println(err)
		return err
	}

	file, err := os.Open("test_images/bcs_640x480.jpg")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	writer.Close()
	req, err := http.NewRequest(method, urlPath, bytes.NewReader(body.Bytes()))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(requestDump))

	return nil
}
