// Элементы равностороннего треугольника пронумерованы следующим
// образом: 1 — сторона a , 2 — радиус R1 вписанной окружности ( R1 =
// = a*math.Sqrt(3) / 6 ), 3 — радиус R2 описанной окружности ( R2 = 2*R1 ), 4 — площадь
// S = a^2*math.Sqrt(3)/ 4 . Дан номер одного из этих элементов и его значение. Вывести
// значения остальных элементов данного треугольника (в том же порядке)
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		x int
		y float64
	)

	for x < 1 || x > 4 {
		x = ioutil.Integer("Номер")
	}

	switch {
	case x == 1: // сторона
		fmt.Println("сторона")
		y = ioutil.Number("Значение")
		r1 := (y * math.Sqrt(3)) / 6
		r2 := 2 * r1
		s := (math.Pow(y, 2) * math.Sqrt(3)) / 4
		fmt.Printf("сторона: %v, радиус R1: %v, радиус R2: %v, площадь: %v\n", y, r1, r2, s)
	case x == 2: // радиус R1
		fmt.Println("радиус R1")
		r1 := ioutil.Number("Значение")
		y = (r1 * 6) / math.Sqrt(3)
		r2 := 2 * r1
		s := (math.Pow(y, 2) * math.Sqrt(3)) / 4
		fmt.Printf("сторона: %v, радиус R1: %v, радиус R2: %v, площадь: %v\n", y, r1, r2, s)
	case x == 3: // радиус R2
		fmt.Println("радиус R2")
		r2 := ioutil.Number("Значение")
		r1 := r2 / 2
		y = (r1 * 6) / math.Sqrt(3)
		s := (math.Pow(y, 2) * math.Sqrt(3)) / 4
		fmt.Printf("сторона: %v, радиус R1: %v, радиус R2: %v, площадь: %v\n", y, r1, r2, s)
	case x == 4: // площадь
		fmt.Println("площадь")
		s := ioutil.Number("Значение")
		y = math.Sqrt(((s * 4) / math.Sqrt(3)))
		r1 := (y * math.Sqrt(3)) / 6
		r2 := 2 * r1
		fmt.Printf("сторона: %v, радиус R1: %v, радиус R2: %v, площадь: %v\n", y, r1, r2, s)
	}
}
