// Даны два ненулевых числа. Найти сумму, разность, произведение и
// частное их модулей
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	fmt.Println("Необходимо найти сумму, разность, произведение и частное квадратов чисел")
	x = ioutil.NotNullNumber("Введите число 1")
	x = ioutil.ModNumber(x)
	y = ioutil.NotNullNumber("Введите число 2")
	y = ioutil.ModNumber(y)
	fmt.Println("Сумма:\t\t", x+y)
	fmt.Println("Разность:\t", x-y)
	fmt.Println("Произведение:\t", x*y)
	fmt.Println("Частное:\t", x/y)
}
