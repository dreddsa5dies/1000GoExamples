package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests1 = []struct {
		inputA int
		inputB int
		want1  int
		want2  int
	}{
		{20, 3, 2, 6},
	}
	for _, test := range tests1 {
		if got1, got2 := countLoop(test.inputA, test.inputB); got1 != test.want1 || got2 != test.want2 {
			t.Errorf("countLoop(%v, %v) = %v, %v", test.inputA, test.inputB, got1, got2)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(8, 3)
	}
}
