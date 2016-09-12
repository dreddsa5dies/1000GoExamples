// Даны числа x , y . Проверить истинность высказывания: «Точка с ко-
// ординатами ( x , y ) лежит в четвертой координатной четверти»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	var z bool
	x = ioutil.Number("число x")
	y = ioutil.Number("число y")
	z = (y < 0) && (x > 0)
	fmt.Println(z)
}
