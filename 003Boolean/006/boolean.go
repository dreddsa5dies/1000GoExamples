// Даны три целых числа: A , B , C . Проверить истинность высказыва-
// ния: «Справедливо двойное неравенство A < B < C »
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
	fmt.Println(a < b && b < c)
}
