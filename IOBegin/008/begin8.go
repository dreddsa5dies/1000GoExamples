// Даны два числа a и b. Найти их среднее арифметическое: (a + b)/2
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x, y float64
	fmt.Println("Необходимо найти среднее арифметическое чисел")
	x = util.Number("Введите число 1")
	y = util.Number("Введите число 2")
	fmt.Println("Среднее арифметическое чисел:\t", (x+y)/2)
}
