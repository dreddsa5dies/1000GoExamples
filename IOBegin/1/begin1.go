// Дана сторона квадрата a. Найти его периметр P = 4·a.
package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Необходимо найти периметр квадрата")
	x = number("Введите значение стороны")
    fmt.Println("Периметр равен:\t", 4*x)
}

func number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	fmt.Scan(&num)
	return num
}
