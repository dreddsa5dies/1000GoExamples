package main

import "testing"

func TestFooN(t *testing.T) {
	var tests1 = []struct {
		inputN uint
		inputK uint
		want   float64
	}{
		{3, 2, 14},
	}
	for _, test := range tests1 {
		if got := foo1(test.inputN, test.inputK); got != test.want {
			t.Errorf("foo1(%v, %v) = %v", test.inputN, test.inputK, got)
		}
	}
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo1(3, 2)
	}
}
