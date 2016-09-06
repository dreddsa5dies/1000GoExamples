// Даны целочисленные координаты трех вершин прямоугольника, стороны
// которого параллельны координатным осям. Найти координаты его четвер-
// той вершины
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x1, y1, x2, y3, x4, y4 int
	x1 = util.Integer("X1")
	y1 = util.Integer("Y1")

	x2 = x1
	for x1 == x2 {
		x2 = util.Integer("X2")
	}
	fmt.Printf("Y2: %v\n", y1)

	fmt.Printf("X3: %v\n", x2)
	y3 = y1
	for y3 == y1 {
		y3 = util.Integer("Y3")
	}

	if y3 < 0 {
		x4 = x1
	} else {
		x4 = -x1
	}

	if x1 < 0 {
		y4 = -y3
	} else {
		y4 = y3
	}

	fmt.Printf("X4: %v\tY4: %v\n", x4, y4)
}
