package main

import "testing"

func TestPrintInt(t *testing.T) {
	var v string
	v = printChinaYear(1984)
	if v != "год зеленой крысы" {
		t.Error(`printChinaYear(1984) = false`)
	}
	v = printChinaYear(1986)
	if v != "год зеленого тигра" {
		t.Error(`printChinaYear(1984) = false`)
	}
}
