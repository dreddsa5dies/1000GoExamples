// Даны два целых числа: A , B . Проверить истинность высказывания:
// «Справедливы неравенства A ≥ 0 или B < –2»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a, b int
	a = ioutil.Integer("целое число A")
	b = ioutil.Integer("целое число B")
	fmt.Println(a >= 0 || b < -2)
}
