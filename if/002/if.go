// Дано целое число. Если оно является положительным, то прибавить к не-
// му 1; в противном случае вычесть из него 2. Вывести полученное число
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
	} else {
		x = x - 2
	}
	fmt.Println(x)
}
