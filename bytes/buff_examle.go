package main

import (
	"bytes"
	"fmt"
	"os"
)

func main()  {
	file, _ := os.Open("./test")
	buf := bytes.NewBufferString("hello world")
	buf.ReadFrom(file)
	fmt.Println(buf.String())
}