package main

import "testing"

var n int = 100000

func BenchmarkWithSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[string]int, n)
		for j := 0; j < n; j++ {
			m["hhs"+string(rune(j))] = j
		}
	}
}

func BenchmarkWithoutSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[string]int)
		for j := 0; j < n; j++ {
			m["hhs"+string(rune(j))] = j
		}
	}
}
