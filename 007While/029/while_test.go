package main

import "testing"

func TestCountLoop(t *testing.T) {
	var tests = []struct {
		input                  float64
		wantK, wantAK, wantAK1 float64
	}{
		{12, 2, 2, 1},
	}
	for _, test := range tests {
		if gotK, gotAK, gotAK1 := countLoop(test.input); gotK != test.wantK && gotAK != test.wantAK && gotAK1 != test.wantAK1 {
			t.Errorf("countLoop(%v) = %v, %v, %v", test.input, gotK, gotAK, gotAK1)
		}
	}
}

func BenchmarkCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countLoop(18)
	}
}
