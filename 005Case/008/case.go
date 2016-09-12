// Даны два целых числа: D (день) и M (месяц), определяющие правиль-
// ную дату невисокосного года. Вывести значения D и M для даты, предше-
// ствующей указанной
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
	case y == 1:
		if x == 1 {
			fmt.Printf("%v (день) и %v (месяц)\n", 31, 12)
		} else {
			switch {
			case x == 5 || x == 7 || x == 8 || x == 10 || x == 12:
				fmt.Printf("%v (день) и %v (месяц)\n", 30, x-1)
			case x == 4 || x == 6 || x == 9 || x == 11:
				fmt.Printf("%v (день) и %v (месяц)\n", 31, x-1)
			case x == 3:
				fmt.Printf("%v (день) и %v (месяц)\n", 28, x-1)
			}
		}
	case y > 1:
		fmt.Printf("%v (день) и %v (месяц)\n", y-1, x)
	}
}
