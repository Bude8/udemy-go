package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Implementing a custom writer
type logWriter struct{}

func main() {
	// resp Body is type io.ReadCloser which is an interface for Reader, Closer
	// type ReadCloser interface {
	// 	Reader
	// 	Closer
	// }
	// type Reader interface {
	//     Read(p []byte) (n int, err error)
	// }
	// type Closer interface {
	//     Close() error
	// }

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// // Thing to read data into
	// // the Read function cannot automatically resize slices, so we start with a large size
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Source of data -> Reader -> []byte (output data that anyone can work with)
	// []byte -> Writer -> some form of output

	// something that implements the Writer interface
	// os.Stdout -> value of type File -> File has a function called 'Write' -> Therefore, it implements the 'Write' interface

	// something that implemments the Reader interface
	// resp.Body
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many btes:", len(bs))
	return len(bs), nil
}
