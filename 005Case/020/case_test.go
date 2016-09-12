package main

import "testing"

func TestPrintZodiac(t *testing.T) {
	var tests = []struct {
		inputM int
		inputD int
		want   string
	}{
		{1, 21, "Водолей"},
		{2, 20, "Рыбы"},
		{3, 22, "Овен"},
		{4, 24, "Телец"},
	}
	for _, test := range tests {
		if got := printZodiac(test.inputM, test.inputD); got != test.want {
			t.Errorf("printZodiac(%v, %v) = %v", test.inputD, test.inputM, got)
		}
	}
}
