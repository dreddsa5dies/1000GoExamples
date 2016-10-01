package main

import "testing"

func TestSeries(t *testing.T) {
	var tests = []struct {
		inputR float64
		want   int
	}{
		{1.1, 1},
	}
	for _, test := range tests {
		if got := series(test.inputR); got != test.want {
			t.Errorf("series(%v) = %v", test.inputR, got)
		}
	}
}

func BenchmarkSeries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		series(12.2)
	}
}
