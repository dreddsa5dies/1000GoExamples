package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		inputA int
		inputB int
		want   int
	}{
		{10, 2, 2},
		{3, 1, 1},
		{18, 12, 6},
	}
	for _, test := range tests {
		if got := countLoop(test.inputA, test.inputB); got != test.want {
			t.Errorf("countLoop(%v, %v) = %v", test.inputA, test.inputB, got)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(18, 12)
	}
}
