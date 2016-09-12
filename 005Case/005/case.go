// Арифметические действия над числами пронумерованы следующим
// образом: 1 — сложение, 2 — вычитание, 3 — умножение, 4 — деление.
// Дан номер действия N (целое число в диапазоне 1–4) и вещественные чис-
// ла A и B ( В не равно 0). Выполнить над числами указанное действие и вы-
// вести результат
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, z uint
	for x < 1 || x > 4 {
		x = ioutil.UInteger("действие")
	}

	z = ioutil.UInteger("A")
	for y == 0 {
		y = ioutil.UInteger("B")
	}

	switch {
	case x == 1:
		fmt.Printf("%v\n", z+y)
	case x == 2:
		fmt.Printf("%v\n", z-y)
	case x == 3:
		fmt.Printf("%v\n", z*y)
	case x == 4:
		fmt.Printf("%v\n", z/y)
	}
}
