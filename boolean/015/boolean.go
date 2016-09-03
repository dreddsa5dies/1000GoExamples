//Даны три целых числа: A , B , C . Проверить истинность высказыва-
// ния: «Ровно два из чисел A , B , C являются положительными»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a, b, c int
	a = util.Integer("целое число A")
	b = util.Integer("целое число B")
	c = util.Integer("целое число C")
	x := (a > 0 && b > 0 && c < 0)
	y := (a < 0 && b > 0 && c > 0)
	z := (a > 0 && b < 0 && c > 0)
	fmt.Println(x || y || z)
}
