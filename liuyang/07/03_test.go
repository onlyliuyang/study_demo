package main

import (
	"testing"
	"unsafe"
)

func BenchmarkByteString(b *testing.B) {
	str := "this is a string"
	var bs []byte

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs = []byte(str)
		str = string(bs)
	}
	b.StopTimer()
}

func Sbyte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Str2sByte(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s                                                  //把s的地址赋给b
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s) //修改容量为长度
	return
}

func BenchmarkByteString2(b *testing.B) {
	str := "this is a string"
	var bs []byte

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs = Str2sByte(str)
		str = Sbyte2Str(bs)
	}
	b.StopTimer()
}
