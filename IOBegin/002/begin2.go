// Дана сторона квадрата a. Найти его площадь S = a^2 .
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x float64
	fmt.Println("Необходимо найти площадь квадрата")
	x = util.Number("Введите значение стороны")
	fmt.Println("Площадь равна:\t", x*x)
}
