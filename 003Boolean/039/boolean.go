// Даны координаты двух различных полей шахматной доски x 1 , y 1 ,
// x 2 , y 2 (целые числа, лежащие в диапазоне 1–8). Проверить истинность вы-
// сказывания: «Ферзь за один ход может перейти с одного поля на другое»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x1, x2, y1, y2 int
	var a, b bool
	for x1 < 1 || x1 > 8 {
		x1 = ioutil.Integer("число x1")
	}
	for y1 < 1 || y1 > 8 {
		y1 = ioutil.Integer("число y1")
	}
	for x2 < 1 || x2 > 8 {
		x2 = ioutil.Integer("число x2")
	}
	for y2 < 1 || y2 > 8 {
		y2 = ioutil.Integer("число y2")
	}
	a = ioutil.ModNumber(float64(x2-x1)) == ioutil.ModNumber(float64(y2-y1))
	b = (x1 == x2) || (y1 == y2)
	fmt.Println(a || b)
}
