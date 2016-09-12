// Даны два неотрицательных числа a и b. Найти их среднее геометрическое,
// то есть квадратный корень из их произведения
package main

import (
	"fmt"
    "math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	fmt.Println("Необходимо найти среднее геометрическое чисел")
	x = ioutil.NoNNumber("Введите число 1")
	y = ioutil.NoNNumber("Введите число 2")
	fmt.Println("Среднее геометрическое чисел:\t", math.Sqrt(x*y))
}
