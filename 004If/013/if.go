// Даны три числа. Найти среднее из них (то есть число, расположенное ме-
// жду наименьшим и наибольшим)
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
		if y < z {
			fmt.Printf("среднее: %v\n", y)
		} else {
			fmt.Printf("среднее: %v\n", z)
		}
	} else if y < x && y < z {
		if x < z {
			fmt.Printf("среднее: %v\n", x)
		} else {
			fmt.Printf("среднее: %v\n", z)
		}
	} else {
		if x < y {
			fmt.Printf("среднее: %v\n", x)
		} else {
			fmt.Printf("среднее: %v\n", y)
		}
	}
}
