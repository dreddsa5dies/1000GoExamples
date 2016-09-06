// Даны целочисленные координаты трех вершин прямоугольника, стороны
// которого параллельны координатным осям. Найти координаты его четвер-
// той вершины
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x1, y1, x2, y2, x3, y3 int
	x1 = util.Integer("X1")
	y1 = util.Integer("Y1")

	x2 = util.Integer("X2")
	y2 = util.Integer("Y2")

	x3 = util.Integer("X3")
	y3 = util.Integer("Y4")

	if x1 > 0 && y1 > 0 {
		fmt.Printf(": %v\n", 1)
	}
}
