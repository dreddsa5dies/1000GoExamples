// Даны координаты поля шахматной доски x , y (целые числа, лежа-
// щие в диапазоне 1–8). Учитывая, что левое нижнее поле доски (1, 1) явля-
// ется черным, проверить истинность высказывания: «Данное поле является
// белым»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y uint
	var a bool
	for x < 1 || x > 8 {
		x = ioutil.UInteger("число x")
	}
	for y < 1 || y > 8 {
		y = ioutil.UInteger("число y")
	}
	a = ((x%2 == 0) && (y%2 == 0)) || ((x%2 != 0) && (y%2 != 0))
	fmt.Println(!a)
}
