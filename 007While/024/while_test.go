package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{2, true},
		{3, true},
		{4, false},
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
