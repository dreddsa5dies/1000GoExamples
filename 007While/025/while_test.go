package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{2, 3},
		{5, 8},
	}
	for _, test := range tests {
		if got := countLoop(test.input); got != test.want {
			t.Errorf("countLoop(%v) = %v", test.input, got)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(18)
	}
}
