// Даны целые числа a , b , c , являющиеся сторонами некоторого тре-
// угольника. Проверить истинность высказывания: «Треугольник со сторо-
// нами a , b , c является равнобедренным»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b, c int
	var x bool
	a = util.Integer("число a")
	b = util.Integer("число b")
	c = util.Integer("число c")
	x = a == b || a == c || b == c
	fmt.Println(x)
}
