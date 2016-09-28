package main

import "testing"

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(18)
	}
}
