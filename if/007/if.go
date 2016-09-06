// Даны два числа. Вывести порядковый номер меньшего из них
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y int
	x = util.Integer("число 1")
	y = util.Integer("число 2")
	if x > y {
		x = 2
		fmt.Printf("номер: %v\n", x)
	} else {
		y = 1
		fmt.Printf("номер: %v\n", y)
	}
}
