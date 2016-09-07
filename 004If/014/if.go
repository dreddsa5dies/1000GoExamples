// Даны три числа. Вывести вначале наименьшее, а затем наибольшее из дан-
// ных чисел
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y, z int
	x = util.Integer("число x")
	y = util.Integer("число y")
	z = util.Integer("число z")
	if x < y && x < z {
		fmt.Printf("наименьшее: %v\n", x)
		if y > z {
			fmt.Printf("наибольшее: %v\n", y)
		} else {
			fmt.Printf("наибольшее: %v\n", z)
		}
	} else if y < x && y < z {
		fmt.Printf("наименьшее: %v\n", y)
		if x > z {
			fmt.Printf("наибольшее: %v\n", x)
		} else {
			fmt.Printf("наибольшее: %v\n", z)
		}
	} else {
		fmt.Printf("наименьшее: %v\n", z)
		if x > y {
			fmt.Printf("наибольшее: %v\n", x)
		} else {
			fmt.Printf("наибольшее: %v\n", y)
		}
	}
}
