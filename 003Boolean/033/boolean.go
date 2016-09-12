// Даны целые числа a , b , c . Проверить истинность высказывания:
// «Существует треугольник со сторонами a , b , c »
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a, b, c uint
	var x bool
	for a == 0 || b == 0 || c == 0 {
		a = ioutil.UInteger("число a")
		b = ioutil.UInteger("число b")
		c = ioutil.UInteger("число c")
	}
	x = (a+b > c) && (a+c > b) && (c+b > a)
	fmt.Println(x)
}
