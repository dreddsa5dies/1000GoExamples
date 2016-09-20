package main

import "testing"

func TestFooN(t *testing.T) {
	var tests1 = []struct {
		inputN int
		inputX float64
		want   float64
	}{
		{5, 0.1, 0.09966866666666667}, // не очень понимаю почему inputN и inputX меняются местами при go test
	}
	for _, test := range tests1 {
		if got := foo1(test.inputN, test.inputX); got != test.want {
			t.Errorf("foo1(%v, %v) = %v", test.inputN, test.inputX, got)
		}
	}
}

func BenchmarkFoo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo1(2, 0.1)
	}
}
