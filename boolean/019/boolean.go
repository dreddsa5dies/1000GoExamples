//Проверить истинность высказывания: «Среди трех данных целых
// чисел есть хотя бы одна пара взаимно противоположных»
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
	x = (a+b == 0) || (a+c == 0) || (c+b == 0)
	fmt.Println(x)
}
