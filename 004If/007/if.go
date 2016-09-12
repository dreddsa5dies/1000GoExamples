// Даны два числа. Вывести порядковый номер меньшего из них
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y int
	x = ioutil.Integer("число 1")
	y = ioutil.Integer("число 2")
	if x > y {
		x = 2
		fmt.Printf("номер: %v\n", x)
	} else {
		y = 1
		fmt.Printf("номер: %v\n", y)
	}
}
