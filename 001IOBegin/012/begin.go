// Даны катеты прямоугольного треугольника a и b. Найти его гипотенузу c и периметр P :
// c = sqrt(a^2 + b^2) ,
// P = a + b + c
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var x, y float64
	fmt.Println("Необходимо найти гипотенузу и периметр прямоугольного треугольника")
	x = ioutil.NotNullNumber("Введите катет 1")
	x = ioutil.ModNumber(x)
	y = ioutil.NotNullNumber("Введите катет 2")
	y = ioutil.ModNumber(y)
	z := math.Sqrt(x*x + y*y)
	fmt.Println("Гипотенуза:\t", z)
	fmt.Println("Периметр:\t", x+y+z)
}
