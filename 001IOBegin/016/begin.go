// Найти расстояние между двумя точками с заданными координатами
// x 1 и x 2 на числовой оси: | x 2 – x 1 |
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	x = ioutil.Number("введите координату 1")
	y = ioutil.Number("введите координату 1")
	fmt.Println("расстояние:\t", ioutil.ModNumber(x-y))
}
