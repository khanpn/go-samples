package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Read filename from program arguments. ie. Run the program by using command `go run main.go myTextFile.txt`
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}