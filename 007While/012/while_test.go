package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		inputN  int
		wantK   int
		wantSum int
	}{
		{3, 3, 2},
		{7, 6, 3},
		{9, 6, 3},
	}
	for _, test := range tests {
		if gotK, gotSum := countLoop(test.inputN); gotK != test.wantK && gotSum != test.wantSum {
			t.Errorf("countLoop(%v) = %v, %v", test.inputN, gotK, gotSum)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(8)
	}
}
