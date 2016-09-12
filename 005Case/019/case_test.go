package main

import "testing"

func TestPrintInt(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{1984, "год зеленой крысы"},
		{1986, "год зеленого тигра"},
		{1900, "год зеленой крысы"},
	}
	for _, test := range tests {
		if got := printChinaYear(test.input); got != test.want {
			t.Errorf("printChinaYear(%v) = %v", test.input, got)
		}
	}
}
