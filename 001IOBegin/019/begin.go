// Даны координаты двух противоположных вершин прямоугольника:
// ( x 1 , y 1 ), ( x 2 , y 2 ). Стороны прямоугольника параллельны осям координат.
// Найти периметр и площадь данного прямоугольника
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x1, y1, x2, y2 float64
	x1 = util.Number("Введите координату х1")
	y1 = util.Number("Введите координату y1")
	x2 = util.Number("Введите координату х2")
	y2 = util.Number("Введите координату y2")
	fmt.Println("периметр прямоугольника\t: ", 2*(util.ModNumber(x1-x2)+util.ModNumber(y1-y2)))
	fmt.Println("площадь прямоугольника\t: ", util.ModNumber(x1-x2)*util.ModNumber(y1-y2))
}
