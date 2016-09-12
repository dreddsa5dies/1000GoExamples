// Дано трехзначное число. Проверить истинность высказывания:
// «Все цифры данного числа различны»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a int
	var x bool
	for a > 999 || a < 100 {
		a = ioutil.Integer("целое трехзначное число A")
	}
	x = (a%10) != (a/100) && (a%10) != ((a-(100*(a/100)))/10) && (a/100) != (a-(100*(a/100)))
	fmt.Println(x)
}
