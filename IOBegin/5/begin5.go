// Дана длина ребра куба a. Найти объем куба V = a^3 и площадь его поверхности S = 6·a^2
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x float64
	fmt.Println("Необходимо найти объем куба и площадь его поверхности")
	x = util.Number("Введите значение длины ребра куба")
	fmt.Println("Объем равен:\t", x*x*x)
	fmt.Println("Площадь равна:\t", 6*x*x)
}
