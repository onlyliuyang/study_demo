package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func stringToByte(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	runtime.KeepAlive(&s)
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytesToString(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	runtime.KeepAlive(&b)
	return *(*string)(unsafe.Pointer(&sh))
}

func main()  {
	var str string = "Hello World"
	bytes := stringToByte(str)
	fmt.Println(bytes)
	fmt.Println(bytesToString(bytes))

	//uintptr()

}