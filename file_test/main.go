package main

import (
	"fmt"
	"io"
	"os"
)

type fileLogger struct{}

func main() {
	fileArg := os.Args[1]
	file, err := os.Open(fileArg)
	if err != nil {
		fmt.Println(err)
	}

	lf := fileLogger{}
	io.Copy(lf, file)
}

func (f fileLogger) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}
