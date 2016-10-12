// Описать процедуру RectPS( x 1 , y 1 , x 2 , y 2 , P , S ), вычисляющую периметр P
// и площадь S прямоугольника со сторонами, параллельными осям коорди-
// нат, по координатам ( x 1 , y 1 ), ( x 2 , y 2 ) его противоположных вершин ( x 1 , y 1 ,
// x 2 , y 2 — входные, P и S — выходные параметры вещественного типа).
// С помощью этой процедуры найти периметры и площади трех прямо-
// угольников с данными противоположными вершинами

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	arr := ioutil.RandomFlt(ioutil.Integer("Число треугольников"))
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			p, s := rectPS(arr[i])
			fmt.Printf("сторона = %v, периметр = %v, площадь = %v\n", arr[i], p, s)
		}
	}
}

func rectPS(a float64) (p, s float64) {
	p = 3 * a
	s = math.Pow(a, 2) * math.Sqrt(3) / 4
	return p, s
}
