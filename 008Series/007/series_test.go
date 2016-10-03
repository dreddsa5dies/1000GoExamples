package main

import "testing"

func BenchmarkSeries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		series(12)
	}
}
