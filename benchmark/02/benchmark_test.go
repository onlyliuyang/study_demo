package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	bytes := []byte("I am byte array!")
	str := string(bytes)
	bytes[0] = 'i'
	fmt.Println(str)

	str1 := (*string)(unsafe.Pointer(&bytes))
	bytes[0] = 'W'
	fmt.Println(*str1)
}

