// Даны координаты двух различных полей шахматной доски x 1 , y 1 ,
// x 2 , y 2 (целые числа, лежащие в диапазоне 1–8). Проверить истинность вы-
// сказывания: «Король за один ход может перейти с одного поля на другое»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x1, x2, y1, y2 uint
	var a bool
	for x1 < 1 || x1 > 8 {
		x1 = ioutil.UInteger("число x1")
	}
	for y1 < 1 || y1 > 8 {
		y1 = ioutil.UInteger("число y1")
	}
	for x2 < 1 || x2 > 8 {
		x2 = ioutil.UInteger("число x2")
	}
	for y2 < 1 || y2 > 8 {
		y2 = ioutil.UInteger("число y2")
	}
	a = ((x2 == (x1 - 1)) || (x2 == (x1 + 1))) || (y2 == (y1 - 1)) || (y2 == (y1 + 1))
	fmt.Println(a)
}
