// Найти расстояние между двумя точками с заданными координатами
// ( x 1 , y 1 ) и ( x 2 , y 2 ) на плоскости. Расстояние вычисляется по формуле
// sqrt(( x 2 − x 1 )^2 + ( y 2 − y 1 )^2)
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x1, y1, x2, y2 float64
	x1 = ioutil.Number("Введите координату х1")
	y1 = ioutil.Number("Введите координату y1")
	x2 = ioutil.Number("Введите координату х2")
	y2 = ioutil.Number("Введите координату y2")
	z1 := x2 - x1
	z2 := y2 - y1
	fmt.Println("расстояние между двумя точками\t: ", math.Sqrt(z1*z1+z2*z2))
}
