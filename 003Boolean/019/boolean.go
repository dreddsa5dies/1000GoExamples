//Проверить истинность высказывания: «Среди трех данных целых
// чисел есть хотя бы одна пара взаимно противоположных»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a, b, c int
	var x bool
	a = ioutil.Integer("целое число A")
	b = ioutil.Integer("целое число B")
	c = ioutil.Integer("целое число C")
	x = (a+b == 0) || (a+c == 0) || (c+b == 0)
	fmt.Println(x)
}
