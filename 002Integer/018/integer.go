// Дано целое число, большее 999. Используя одну операцию деления
// нацело и одну операцию взятия остатка от деления, найти цифру, соответ-
// ствующую разряду тысяч в записи этого числа
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x uint
	for x < 999 {
		x = ioutil.UInteger("Введите целое число, большее 999")
	}
	x = x / 1000
	x = x % 10
	fmt.Printf("разряд тысяч: %v\n", x)
}
