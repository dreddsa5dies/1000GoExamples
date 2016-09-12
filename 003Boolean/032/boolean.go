// Даны целые числа a , b , c , являющиеся сторонами некоторого тре-
// угольника. Проверить истинность высказывания: «Треугольник со сторо-
// нами a , b , c является прямоугольным»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a, b, c uint
	var x bool
	a = ioutil.UInteger("число a")
	b = ioutil.UInteger("число b")
	c = ioutil.UInteger("число c")
	x = (c*c == (a*a + b*b)) || (a*a == (c*c + b*b)) || (b*b == (a*a + c*c))
	fmt.Println(x)
}
