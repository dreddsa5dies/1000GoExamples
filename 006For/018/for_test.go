package main

import "testing"

func TestFoo1(t *testing.T) {
	var tests = []struct {
		inputA float64
		inputN int
		want   float64
	}{
		{2, 4, 11},
	}
	for _, test := range tests {
		if got := foo1(test.inputA, test.inputN); got != test.want {
			t.Errorf("foo1(%v, %v) = %v", test.inputA, test.inputN, got)
		}
	}
}
