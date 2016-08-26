// Даны два ненулевых числа. Найти сумму, разность, произведение и
// частное их модулей
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y float64
	fmt.Println("Необходимо найти сумму, разность, произведение и частное квадратов чисел")
	x = util.NotNullNumber("Введите число 1")
	x = util.ModNumber(x)
	y = util.NotNullNumber("Введите число 2")
	y = util.ModNumber(y)
	fmt.Println("Сумма:\t\t", x+y)
	fmt.Println("Разность:\t", x-y)
	fmt.Println("Произведение:\t", x*y)
	fmt.Println("Частное:\t", x/y)
}
