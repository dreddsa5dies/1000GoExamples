// Дано трехзначное число. Используя одну операцию деления нацело,
// вывести первую цифру данного числа (сотни)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x uint
	for x < 100 || x > 1000 {
		x = util.UInteger("Введите трехзначное число A")
	}
	fmt.Printf("первая цифра данного числа (сотни): %v\n", x/100)
}
