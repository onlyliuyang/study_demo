package main

import (
	"reflect"
	"testing"
	"unsafe"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func s2b (s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}


var s = "adsfasdfadsfadsfasdfadfadfasdfasdfadsfasdfasdfasdfsadfas"

func BenchmarkB2sFast(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		s2b(s)
	}
}

func BenchmarkB2sStd(b *testing.B)  {
	var _ []byte
	for i:=0; i<b.N; i++ {
		_ = []byte(s)
	}
}


var bt = []byte("adsfasdfadsfadsfasdfadfadfasdfasdfadsfasdfasdfasdfsadfas")

func BenchmarkS2bFast(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		b2s(bt)
	}
}

func BenchmarkS2bStd(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		_ = string(bt)
	}
}