package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type CustomWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")

	if (err != nil) {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(CustomWriter{}, resp.Body)
}

func (CustomWriter) Write(bytes []byte) (int, error) {
	fmt.Println(string(bytes))
	return len(bytes), nil
}