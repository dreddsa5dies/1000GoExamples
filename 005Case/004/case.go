// Дан номер месяца — целое число в диапазоне 1–12 (1 — январь, 2 —
// февраль и т. д.). Определить количество дней в этом месяце для невисо-
// косного года
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x int
	for x < 1 || x > 12 {
		x = ioutil.Integer("число")
	}

	switch {
	case x == 1 || x == 3 || x == 5 || x == 7 || x == 8 || x == 10 || x == 12:
		fmt.Println("31")
	case x == 4 || x == 6 || x == 9 || x == 11:
		fmt.Println("30")
	case x == 2:
		fmt.Println("28")
	}
}
