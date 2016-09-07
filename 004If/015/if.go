// Даны три числа. Найти сумму двух наибольших из них
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
		fmt.Printf("сумма двух наибольших из них: %v\n", y+z)
	} else if y < x && y < z {
		fmt.Printf("сумма двух наибольших из них: %v\n", x+z)
	} else {
		fmt.Printf("сумма двух наибольших из них: %v\n", y+x)
	}
}
