package main

import (
	"fmt"
	"io"
	"os"
)

// Reads the contents of a text file then prints its contents to terminal
// File to open should be provided as an arg

// To open a file, check documentation for 'Open' function in 'os' package
// https://golang.org/pkg/os/#Open

func main() {
	file := os.Args[1]
	f, err := os.Open(file)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
