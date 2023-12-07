package main

import (
	"bytes"
	"fmt"
)

func main()  {
	buf1 := bytes.NewBufferString("hello")
	fmt.Println(buf1)

	buf2 := bytes.NewBuffer([]byte("hello"))
	fmt.Println(buf2)

	buf3 := bytes.NewBufferString("")
	buf4 := bytes.NewBuffer([]byte(""))
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println(buf3, buf4, buf5)
}
