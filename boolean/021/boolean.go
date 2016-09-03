// Дано трехзначное число. Проверить истинность высказывания:
// «Цифры данного числа образуют возрастающую последовательность»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var a int
	var x bool
	for a > 999 || a < 100 {
		a = util.Integer("целое трехзначное число A")
	}
	x = (a/100) < ((a-(100*(a/100)))/10) && ((a-(100*(a/100)))/10) < (a%10)
	fmt.Println(x)
}
