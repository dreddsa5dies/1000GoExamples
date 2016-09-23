package main

import "testing"

func TestFooN(t *testing.T) {
	var tests1 = []struct {
		inputN int
		want   float64
	}{
		{3, 32},
	}
	for _, test := range tests1 {
		if got := foo1(test.inputN); got != test.want {
			t.Errorf("foo1(%v) = %v", test.inputN, got)
		}
	}
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo1(3)
	}
}
