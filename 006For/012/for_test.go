package main

import "testing"

func TestFoo1(t *testing.T) {
	var tests = []struct {
		inputM int
		want   float64
	}{
		{2, 1.3200000000000003},
		{1, 1.1},
	}
	for _, test := range tests {
		if got := foo1(test.inputM); got != test.want {
			t.Errorf("foo1(%v) = %v", test.inputM, got)
		}
	}
}
