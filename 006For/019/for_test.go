package main

import "testing"

func TestFoo1(t *testing.T) {
	var tests = []struct {
		inputN int
		want   float64
	}{
		{4, 24},
	}
	for _, test := range tests {
		if got := foo1(test.inputN); got != test.want {
			t.Errorf("foo1(%v) = %v", test.inputN, got)
		}
	}
}
