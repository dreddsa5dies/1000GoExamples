// Даны два числа. Вывести большее из них
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y int
	x = util.Integer("число x")
	y = util.Integer("число y")
	if x > y {
		fmt.Printf("большее число: %v\n", x)
	} else {
		fmt.Printf("большее число: %v\n", y)
	}
}
