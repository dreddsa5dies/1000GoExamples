package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests1 = []struct {
		inputA int
		inputB int
		want   int
	}{
		{8, 3, 2},
	}
	for _, test := range tests1 {
		if got := countLoop(test.inputA, test.inputB); got != test.want {
			t.Errorf("countLoop(%v, %v) = %v", test.inputA, test.inputB, got)
		}
	}
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(8, 3)
	}
}
