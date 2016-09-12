// Даны числа x , y , x 1 , y 1 , x 2 , y 2 . Проверить истинность высказывания:
// «Точка с координатами ( x , y ) лежит внутри прямоугольника, левая верхняя
// вершина которого имеет координаты ( x 1 , y 1 ), правая нижняя — ( x 2 , y 2 ),
// а стороны параллельны координатным осям»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y, x1, x2, y1, y2 float64
	var a, b bool
	x = ioutil.Number("число x")
	y = ioutil.Number("число y")
	x1 = ioutil.Number("число x1")
	y1 = ioutil.Number("число y1")
	x2 = ioutil.Number("число x2")
	y2 = ioutil.Number("число y2")
	a = ((x < x1) || (x < x2)) && ((x > x1) || (x > x2))
	b = ((y < y1) || (y < y2)) && ((y > y1) || (y > y2))
	fmt.Println(a && b)
}
