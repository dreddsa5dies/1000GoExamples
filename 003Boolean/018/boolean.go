//Проверить истинность высказывания: «Среди трех данных целых
// чисел есть хотя бы одна пара совпадающих»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b, c int
	var x bool
	a = util.Integer("целое число A")
	b = util.Integer("целое число B")
	c = util.Integer("целое число C")
	x = (a == b && a != c && b != c) || (a == c && a != b && b != a) || (c == b && a != c && b != a)
	fmt.Println(x)
}
