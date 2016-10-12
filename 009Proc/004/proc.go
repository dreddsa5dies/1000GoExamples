// Описать процедуру TrianglePS( a , P , S ), вычисляющую по стороне a
// равностороннего треугольника его периметр P = 3· a и площадь S = a^2* Sqrt(3) / 4
// ( a — входной, P и S — выходные параметры; все параметры являются ве-
// щественными). С помощью этой процедуры найти периметры и площади
// трех равносторонних треугольников с данными сторонами

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
			p, s := trianglePS(arr[i])
			fmt.Printf("сторона = %v, периметр = %v, площадь = %v\n", arr[i], p, s)
		}
	}
}

func trianglePS(a float64) (p, s float64) {
	p = 3 * a
	s = math.Pow(a, 2) * math.Sqrt(3) / 4
	return p, s
}
