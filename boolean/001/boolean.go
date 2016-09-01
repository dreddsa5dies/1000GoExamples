// Дано целое число A . Проверить истинность высказывания: «Число
// A является положительным»
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x int
	x = util.Integer("целое число A")
	if x < 0 {
		fmt.Printf("Число %v является отрицательным\n", x)
	} else {
		fmt.Printf("Число %v является положительным\n", x)
	}
}
