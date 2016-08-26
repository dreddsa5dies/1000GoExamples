// Даны длины ребер a, b, c прямоугольного параллелепипеда. Найти его объем V = a·b·c и
// площадь поверхности S = 2·(a·b + b·c + a·c)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b, c float64
	fmt.Println("Необходимо найти объем прямоугольного параллелепипеда и площадь его поверхности")
	a = util.Number("Введите значение длины 1 ребра куба")
	b = util.Number("Введите значение длины 2 ребра куба")
	c = util.Number("Введите значение длины 3 ребра куба")
	fmt.Println("Объем равен:\t", a*b*c)
	fmt.Println("Площадь равна:\t", 2*(a*b+b*c+a*c))
}
