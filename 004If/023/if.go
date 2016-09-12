// Даны целочисленные координаты трех вершин прямоугольника, стороны
// которого параллельны координатным осям. Найти координаты его четвер-
// той вершины
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x1, y1, x2, y3, x4, y4 int
	x1 = ioutil.Integer("X1")
	y1 = ioutil.Integer("Y1")

	x2 = x1
	for x1 == x2 {
		x2 = ioutil.Integer("X2")
	}
	fmt.Printf("Y2: %v\n", y1) // y2=y1

	fmt.Printf("X3: %v\n", x2) // x3=x2
	y3 = y1
	for y3 == y1 {
		y3 = ioutil.Integer("Y3")
	}

	x4 = x1
	y4 = y3

	fmt.Printf("X4: %v\tY4: %v\n", x4, y4)
}
