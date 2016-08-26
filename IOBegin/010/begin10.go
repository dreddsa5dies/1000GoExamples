// Даны два ненулевых числа. Найти сумму, разность, произведение и
// частное их квадратов
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y float64
	fmt.Println("Необходимо найти сумму, разность, произведение и частное квадратов чисел")
	x = util.NotNullNumber("Введите число 1")
	y = util.NotNullNumber("Введите число 2")
	fmt.Println("Сумма квадратов:\t", ((x * x) + (y * y)))
	fmt.Println("Разность квадратов:\t", ((x * x) - (y * y)))
	fmt.Println("Произведение квадратов:\t", ((x * x) * (y * y)))
	fmt.Println("Частное квадратов:\t", ((x * x) / (y * y)))
}
