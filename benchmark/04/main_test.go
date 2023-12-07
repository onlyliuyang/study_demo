package main

import "testing"

func BenchmarkSumString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumString(base)
	}
	b.StopTimer()
}

func BenchmarkSprintfString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SprintfString(base)
	}
	b.StopTimer()
}

func BenchmarkBuilderString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuilderString(base)
	}
	b.StopTimer()
}

func BenchmarkByteString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ByteString(base)
	}
	b.StopTimer()
}

func BenchmarkByteSliceString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ByteSliceString(base)
	}
	b.StopTimer()
}

func BenchmarkJoinString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JoinString(base)
	}
	b.StopTimer()
}
