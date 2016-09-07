// Даны числа x , y . Проверить истинность высказывания: «Точка с ко-
// ординатами ( x , y ) лежит во второй координатной четверти»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y float64
	var z bool
	x = util.Number("число x")
	y = util.Number("число y")
	z = (y > 0) && (x < 0)
	fmt.Println(z)
}
