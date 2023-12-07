package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := "abcdefg"
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", []byte(a))
}
