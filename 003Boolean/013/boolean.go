// Даны три целых числа: A , B , C . Проверить истинность высказыва-
// ния: «Хотя бы одно из чисел A , B , C положительное»
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
	fmt.Println(a > 0 || b > 0 || c > 0)
}
