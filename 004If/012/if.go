// Даны три числа. Найти наименьшее из них
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, z int
	x = ioutil.Integer("число x")
	y = ioutil.Integer("число y")
	z = ioutil.Integer("число z")
	if x < y && x < z {
		fmt.Printf("наименьшее: %v\n", x)
	} else if y < x && y < z {
		fmt.Printf("наименьшее: %v\n", y)
	} else {
		fmt.Printf("наименьшее: %v\n", z)
	}
}
