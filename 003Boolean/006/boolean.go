// Даны три целых числа: A , B , C . Проверить истинность высказыва-
// ния: «Справедливо двойное неравенство A < B < C »
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a, b, c int
	a = ioutil.Integer("целое число A")
	b = ioutil.Integer("целое число B")
	c = ioutil.Integer("целое число C")
	fmt.Println(a < b && b < c)
}
