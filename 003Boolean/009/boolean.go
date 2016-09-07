// Даны два целых числа: A , B . Проверить истинность высказывания:
// «Хотя бы одно из чисел A и B нечетное»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b int
	a = util.Integer("целое число A")
	b = util.Integer("целое число B")
	fmt.Println(a%2 != 0 || b%2 != 0)
}
