// На числовой оси расположены три точки: A , B , C . Определить, какая из
// двух последних точек ( B или C ) расположена ближе к A , и вывести эту точ-
// ку и ее расстояние от точки A
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y, z int
	x = util.Integer("число A")
	y = util.Integer("число B")
	z = util.Integer("число C")

	if x-y > x-z {
		fmt.Printf("B: %v\n", y)
	} else if x-y == x-z {
		fmt.Printf("B & C: %v & %v\n", y, z)
	} else {
		fmt.Printf("C: %v\n", z)
	}
}
