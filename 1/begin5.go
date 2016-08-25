// Дана длина ребра куба a. Найти объем куба V = a^3 и площадь его поверхности S = 6·a^2
package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Println("Необходимо найти объем куба и площадь его поверхности")
	x = number("Введите значение длины ребра куба")
	fmt.Println("Объем равен:\t", x*x*x)
    fmt.Println("Площадь равна:\t", 6*x*x)
}

func number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	fmt.Scan(&num)
	return num
}
