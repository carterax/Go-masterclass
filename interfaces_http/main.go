package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// type out interface {
// 	Write(p []byte) (int, error)
// }

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	lw := logWriter{}
	// io.Copy implements the write interface
	// @param lw Writer
	// @param resp.Body Reader from external source
	io.Copy(lw, resp.Body)
}

// Write function attached to the struct logWriter now implements the Writer interface
// basically function reusability
func (lw logWriter) Write(b []byte) (int, error) {
	fmt.Println("Here is the length", len(b))
	fmt.Println(string(b))
	return len(b), nil
}
