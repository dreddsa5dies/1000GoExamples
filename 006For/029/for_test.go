package main

import "testing"

func TestFooN(t *testing.T) {
	var tests1 = []struct {
		inputN int
		inputA float64
		inputB float64
		want   float64
	}{
		{3, 2, 5, 1},
	}
	for _, test := range tests1 {
		if got := foo1(test.inputN, test.inputA, test.inputB); got != test.want {
			t.Errorf("foo1(%v, %v,%v) = %v", test.inputN, test.inputA, test.inputB, got)
		}
	}
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo1(3, 2, 5)
	}
}
