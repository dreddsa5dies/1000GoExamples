package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		inputA, inputB, inputC float64
		want                   int
	}{
		{1, 1, 1, 1},
		{2, 1, 1, 2},
	}
	for _, test := range tests {
		if got := countLoop(test.inputA, test.inputB, test.inputC); got != test.want {
			t.Errorf("countLoop(%v, %v, %v) = %v", test.inputA, test.inputB, test.inputC, got)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(18, 12, 1)
	}
}
