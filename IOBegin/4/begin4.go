// Дан диаметр окружности d. Найти ее длину L = π·d. В качестве значения π использовать 3.14
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println("Необходимо найти длину окружности")
	x = number("Введите значение диаметра окружности")
	fmt.Println("Длина равна:\t", x*math.Pi)
}

func number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	fmt.Scan(&num)
	return num
}
