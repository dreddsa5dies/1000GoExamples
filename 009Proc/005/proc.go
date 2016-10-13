// Описать процедуру RectPS( x 1 , y 1 , x 2 , y 2 , P , S ), вычисляющую периметр P
// и площадь S прямоугольника со сторонами, параллельными осям коорди-
// нат, по координатам ( x 1 , y 1 ), ( x 2 , y 2 ) его противоположных вершин ( x 1 , y 1 ,
// x 2 , y 2 — входные, P и S — выходные параметры вещественного типа).
// С помощью этой процедуры найти периметры и площади трех прямо-
// угольников с данными противоположными вершинами

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x1, y1, x2, y2 float64
	for i := 0; i < 3; i++ {
		x1 = ioutil.Number("Введите координату х1")
		y1 = ioutil.Number("Введите координату y1")
		x2 = ioutil.Number("Введите координату х2")
		y2 = ioutil.Number("Введите координату y2")
		p, s := rectPS(x1, y1, x2, y2)
		fmt.Printf("периметр прямоугольника\t: %v\n", p)
		fmt.Printf("площадь прямоугольника\t: %v\n", s)
	}

func rectPS(x1, y1, x2, y2 float64) (p, s float64) {
	p = 2 * (ioutil.ModNumber(x1-x2) + ioutil.ModNumber(y1-y2))
	s = ioutil.ModNumber(x1-x2) * ioutil.ModNumber(y1-y2)
	return p, s
}
