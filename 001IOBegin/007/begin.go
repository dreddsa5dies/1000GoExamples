// Найти длину окружности L и площадь круга S заданного радиуса R:
// L = 2·π·R
// S = π·R^2
// В качестве значения π использовать 3.14
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x float64
	fmt.Println("Необходимо найти длину окружности и площадь круга")
	x = ioutil.Number("Введите значение радиуса")
	fmt.Println("Длина окружности:\t", 2*math.Pi*x)
	fmt.Println("Площадь круга:\t\t", x*math.Pi*math.Pi)
}
