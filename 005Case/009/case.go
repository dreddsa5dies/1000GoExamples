// Даны два целых числа: D (день) и M (месяц), определяющие правиль-
// ную дату невисокосного года. Вывести значения D и M для даты, следую-
// щей за указанной
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		x, y uint
	)

	for x < 1 || x > 12 {
		x = ioutil.UInteger("M (месяц)")
	}

	switch {
	case x == 1 || x == 3 || x == 5 || x == 7 || x == 8 || x == 10 || x == 12:
		for y < 1 || y > 31 {
			y = ioutil.UInteger("D (день)")
		}
	case x == 4 || x == 6 || x == 9 || x == 11:
		for y < 1 || y > 30 {
			y = ioutil.UInteger("D (день)")
		}
	case x == 2:
		for y < 1 || y > 28 {
			y = ioutil.UInteger("D (день)")
		}
	}

	switch {
	case x == 1 || x == 3 || x == 5 || x == 7 || x == 8 || x == 10 || x == 12:
		if y == 31 {
			fmt.Printf("%v (день) и %v (месяц)\n", 1, x+1)
		} else {
			fmt.Printf("%v (день) и %v (месяц)\n", y+1, x)
		}
	case x == 4 || x == 6 || x == 9 || x == 11:
		if y == 30 {
			fmt.Printf("%v (день) и %v (месяц)\n", 1, x+1)
		} else {
			fmt.Printf("%v (день) и %v (месяц)\n", y+1, x)
		}
	case x == 2:
		if y == 28 {
			fmt.Printf("%v (день) и %v (месяц)\n", 1, x+1)
		} else {
			fmt.Printf("%v (день) и %v (месяц)\n", y+1, x)
		}
	}
}
