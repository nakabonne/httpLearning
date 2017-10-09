package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "まいける")
	fileWriter, err := writer.CreateFormFile("thumbnail", "hey.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("hey.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:8080", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
