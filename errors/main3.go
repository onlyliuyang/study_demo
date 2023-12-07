package main

import (
	"fmt"
	"io"
	"os"
)

type byteCounter struct {
}

func (b byteCounter) Write(p []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Println(rw)

	w = new(byteCounter)
	if rx, ok := w.(io.ReadWriter); ok {
		fmt.Println(rx)
	}
}
