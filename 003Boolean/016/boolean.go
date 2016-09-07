//Дано целое положительное число. Проверить истинность высказы-
// вания: «Данное число является четным двузначным»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a int
	var x bool
	a = util.Integer("целое число A")
	x = ((9 < a) && (a < 99)) && (a%2 == 0)
	fmt.Println(x)
}
