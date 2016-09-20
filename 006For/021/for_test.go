package main

import "testing"

func TestFooN(t *testing.T) {
	var tests1 = []struct {
		inputN int
		want   float64
	}{
		{1, 2},
	}
	for _, test := range tests1 {
		if got := foo1(test.inputN); got != test.want {
			t.Errorf("foo1(%v) = %v", test.inputN, got)
		}
	}
	var tests2 = []struct {
		inputN int
		want   float64
	}{
		{1, 1},
	}
	for _, test := range tests2 {
		if got := foo2(test.inputN); got != test.want {
			t.Errorf("foo2(%v) = %v", test.inputN, got)
		}
	}
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo1(4)
	}
}

func BenchmarkFoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo2(4)
	}
}
