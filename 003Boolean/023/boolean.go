// Дано четырехзначное число. Проверить истинность высказывания:
// «Данное число читается одинаково слева направо и справа налево»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a int
	var x bool
	for a > 9999 || a < 1000 {
		a = ioutil.Integer("целое трехзначное число A")
	}
	x = ((a / 1000) == (a % 10)) && (a-(1000*(a/1000)))/100 == (a-(100*(a/100)))/10
	fmt.Println(x)
}
