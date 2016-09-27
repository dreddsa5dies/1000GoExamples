package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		inputA  int
		wantK   float64
		wantSum float64
	}{
		{2, 1.833333333333333, 3},
	}
	for _, test := range tests {
		if gotK, gotSum := countLoop(test.inputA); gotK != test.wantK && gotSum != test.wantSum {
			t.Errorf("countLoop(%v) = %v, %v", test.inputA, gotK, gotSum)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(8)
	}
}
