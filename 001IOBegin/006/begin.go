// Даны длины ребер a, b, c прямоугольного параллелепипеда. Найти его объем V = a·b·c и
// площадь поверхности S = 2·(a·b + b·c + a·c)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a, b, c float64
	fmt.Println("Необходимо найти объем прямоугольного параллелепипеда и площадь его поверхности")
	a = ioutil.Number("Введите значение длины 1 ребра куба")
	b = ioutil.Number("Введите значение длины 2 ребра куба")
	c = ioutil.Number("Введите значение длины 3 ребра куба")
	fmt.Println("Объем равен:\t", a*b*c)
	fmt.Println("Площадь равна:\t", 2*(a*b+b*c+a*c))
}
