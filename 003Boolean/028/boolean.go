// Даны числа x , y . Проверить истинность высказывания: «Точка с ко-
// ординатами ( x , y ) лежит в первой или третьей координатной четверти»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	var a, b bool
	x = ioutil.Number("число x")
	y = ioutil.Number("число y")
	a = (y > 0) && (x > 0)
	b = (y < 0) && (x < 0)
	fmt.Println(a || b)
}
