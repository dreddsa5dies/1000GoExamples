// Даны координаты двух различных полей шахматной доски x 1 , y 1 ,
// x 2 , y 2 (целые числа, лежащие в диапазоне 1–8). Проверить истинность вы-
// сказывания: «Конь за один ход может перейти с одного поля на другое»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x1, x2, y1, y2 int
	var a, b, c, d bool
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
	a = (x2 == (x1 - 1)) && ((y2 == (y1 + 2)) || (y2 == (y1 - 2)))
	b = (x2 == (x1 - 2)) && ((y2 == (y1 + 1)) || (y2 == (y1 - 1)))
	c = (x2 == (x1 + 1)) && ((y2 == (y1 + 2)) || (y2 == (y1 - 2)))
	d = (x2 == (x1 + 2)) && ((y2 == (y1 + 1)) || (y2 == (y1 - 1)))
	fmt.Println(a || b || c || d)
}
