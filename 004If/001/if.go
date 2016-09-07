// Дано целое число. Если оно является положительным, то прибавить к не-
// му 1; в противном случае не изменять его. Вывести полученное число
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x int
	x = util.Integer("число")
	if x > 0 {
		x++
	}
	fmt.Println(x)
}
