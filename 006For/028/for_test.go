package main

import "testing"

func TestFooN(t *testing.T) {
	var tests1 = []struct {
		inputN int
		inputX float64
		want   float64
	}{
		{3, 0.1, 1.0488125000000001},
		{2, 0.1, 1.04875},
	}
	for _, test := range tests1 {
		if got := foo1(test.inputN, test.inputX); got != test.want {
			t.Errorf("foo1(%v, %v) = %v", test.inputN, test.inputX, got)
		}
	}

	var tests2 = []struct {
		inputM float64
		want   float64
	}{
		{3, 48},
		{2, 8},
	}
	for _, test := range tests2 {
		if got := foo2(test.inputM); got != test.want {
			t.Errorf("foo2(%v) = %v", test.inputM, got)
		}
	}
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo1(3, 0.1)
	}
}

func BenchmarkFoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo2(3)
	}
}
