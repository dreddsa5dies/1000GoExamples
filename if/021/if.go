// Даны целочисленные координаты точки на плоскости. Если точка совпа-
// дает с началом координат, то вывести 0. Если точка не совпадает с началом
// координат, но лежит на оси OX или OY , то вывести соответственно 1 или 2.
// Если точка не лежит на координатных осях, то вывести 3
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y int
	x = util.Integer("X")
	y = util.Integer("Y")

	if x == 0 && y == 0 {
		fmt.Printf(": %v\n", y)
	} else {
		if x == 0 {
			fmt.Printf(": %v\n", 1)
		} else if y == 0 {
			fmt.Printf(": %v\n", 2)
		} else {
			fmt.Printf(": %v\n", 3)
		}
	}
}
