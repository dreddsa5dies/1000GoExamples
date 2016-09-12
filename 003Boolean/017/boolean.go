//Дано целое положительное число. Проверить истинность высказы-
// вания: «Данное число является нечетным трехзначным»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a uint
	var x bool
	a = ioutil.UInteger("целое положительное число A")
	x = ((99 < a) && (a < 1000)) && (a%2 != 0)
	fmt.Println(x)
}
