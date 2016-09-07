// Дано целое число. Если оно является положительным, то прибавить к не-
// му 1; если отрицательным, то вычесть из него 2; если нулевым, то заме-
// нить его на 10. Вывести полученное число
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
	} else if x == 0 {
		x = 10
	} else {
		x = x - 2
	}
	fmt.Println(x)
}
