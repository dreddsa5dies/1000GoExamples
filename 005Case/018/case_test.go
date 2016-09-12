package main

import "testing"

func TestPrintInt(t *testing.T) {
	var v string
	v = printInt(111)
	if v != "сто одинадцать " {
		t.Error(`PrintInt(111) = false`)
	}
	v = printInt(205)
	if v != "двести  пять" {
		t.Error(`PrintInt(111) = false`)
	}
	v = printInt(101)
	if v != "сто  один" {
		t.Error(`PrintInt(111) = false`)
	}
}
