package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		input int
		want1 int
		want2 int
	}{
		{2, 1, 3},
		{5, 3, 8},
	}
	for _, test := range tests {
		if got1, got2 := countLoop(test.input); got1 != test.want1 && got2 != test.want2 {
			t.Errorf("countLoop(%v) = %v, %v", test.input, got1, got2)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(18)
	}
}
