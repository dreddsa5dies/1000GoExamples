// Даны координаты точки, не лежащей на координатных осях OX и OY .
// Определить номер координатной четверти, в которой находится данная
// точка
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y int
	for x == 0 || y == 0 {
		x = util.Integer("X")
		y = util.Integer("Y")
	}

	if x > 0 && y > 0 {
		fmt.Printf(": %v\n", 1)
	} else if x < 0 && y > 0 {
		fmt.Printf(": %v\n", 2)
	} else {
		if x < 0 && y < 0 {
			fmt.Printf(": %v\n", 3)
		} else if x > 0 && y < 0 {
			fmt.Printf(": %v\n", 4)
		}
	}
}
