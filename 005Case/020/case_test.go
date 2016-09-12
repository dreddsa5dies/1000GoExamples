package main

import "testing"

func TestPrintInt(t *testing.T) {
	var tests = []struct {
		inputD int
		inputM int
		want   string
	}{
		{21, 1, "Водолей"},
		{20, 2, "Рыбы"},
		{22, 3, "Овен"},
		{24, 4, "Телец"},
	}
	for _, test := range tests {
		if got := printZodiac(test.inputD, test.inputM); got != test.want {
			t.Errorf("printZodiac(%v, %v) = %v", test.inputD, test.inputM, got)
		}
	}
}
