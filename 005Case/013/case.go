// Элементы равнобедренного прямоугольного треугольника пронумеро-
// ваны следующим образом: 1 — катет a , 2 — гипотенуза c = a*math.Sqrt(2) , 3 — вы-
// сота h , опущенная на гипотенузу ( h = c /2), 4 — площадь S = c · h /2. Дан но-
// мер одного из этих элементов и его значение. Вывести значения остальных
// элементов данного треугольника (в том же порядке)
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
	case x == 1: // катет
		fmt.Println("катет")
		y = ioutil.Number("Значение")
		c := y * math.Sqrt(2)
		h := c / 2
		s := c * h / 2
		fmt.Printf("катет: %v, гипотенуза: %v, высота: %v, площадь: %v\n", y, c, h, s)
	case x == 2: // гипотенуза
		fmt.Println("гипотенуза")
		c := ioutil.Number("Значение")
		y = c / math.Sqrt(2)
		h := c / 2
		s := c * h / 2
		fmt.Printf("катет: %v, гипотенуза: %v, высота: %v, площадь: %v\n", y, c, h, s)
	case x == 3: // высота
		fmt.Println("высота")
		h := ioutil.Number("Значение")
		c := h * 2
		y = c / math.Sqrt(2)
		s := c * h / 2
		fmt.Printf("катет: %v, гипотенуза: %v, высота: %v, площадь: %v\n", y, c, h, s)
	case x == 4: // площадь
		fmt.Println("площадь")
		s := ioutil.Number("Значение")
		h := math.Sqrt(s)
		c := 2 * h
		y = c / math.Sqrt(2)
		fmt.Printf("катет: %v, гипотенуза: %v, высота: %v, площадь: %v\n", y, c, h, s)
	}
}
