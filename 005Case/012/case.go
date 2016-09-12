// Элементы окружности пронумерованы следующим образом: 1 — ра-
// диус R , 2 — диаметр D = 2· R , 3 — длина L = 2· π · R , 4 — площадь круга
// S = π · R^2 . Дан номер одного из этих элементов и его значение. Вывести зна-
// чения остальных элементов данной окружности (в том же порядке). В ка-
// честве значения π использовать 3.14
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var (
		x int
		y float64
	)
	const pi = math.Pi

	for x < 1 || x > 4 {
		x = util.Integer("Номер")
	}

	switch {
	case x == 1: // радиус
		fmt.Println("радиус")
		y = util.Number("Значение")
		d := 2 * y
		l := 2 * pi * y
		s := pi * math.Pow(y, 2)
		fmt.Printf("Радиус: %v, диаметр:: %v, длина:: %v, площадь:: %v\n", y, d, l, s)
	case x == 2: // диаметр
		fmt.Println("диаметр")
		d := util.Number("Значение")
		y = d / 2
		l := 2 * pi * y
		s := pi * math.Pow(y, 2)
		fmt.Printf("Радиус: %v, диаметр:: %v, длина:: %v, площадь:: %v\n", y, d, l, s)
	case x == 3: // длина
		fmt.Println("длина")
		l := util.Number("Значение")
		y = l / (2 * pi)
		d := 2 * y
		s := pi * math.Pow(y, 2)
		fmt.Printf("Радиус: %v, диаметр:: %v, длина:: %v, площадь:: %v\n", y, d, l, s)
	case x == 4: // площадь круга
		fmt.Println("площадь круга")
		s := util.Number("Значение")
		y = math.Sqrt(s / pi)
		d := 2 * y
		l := 2 * pi * y
		fmt.Printf("Радиус: %v, диаметр:: %v, длина:: %v, площадь:: %v\n", y, d, l, s)
	}
}
