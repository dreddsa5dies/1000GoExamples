// Дана сторона квадрата a. Найти его площадь S = a^2 .
package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Необходимо найти площадь квадрата")
	x = number("Введите значение стороны")
	fmt.Println("Площадь равна:\t", x*x)
}

func number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	fmt.Scan(&num)
	return num
}
