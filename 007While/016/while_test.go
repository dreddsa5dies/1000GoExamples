package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		inputP     float64
		wantMounth int
		wantSum    float64
	}{
		{10, 32, 211.13776745352556},
	}
	for _, test := range tests {
		if gotSum, gotMounth := countLoop(test.inputP); gotSum != test.wantSum && gotMounth != test.wantMounth {
			t.Errorf("countLoop(%v) = %v, %v", test.inputP, gotSum, gotMounth)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(8)
	}
}
