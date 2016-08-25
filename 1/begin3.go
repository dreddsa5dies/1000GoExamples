// Даны стороны прямоугольника a и b. Найти его площадь S = a·b и периметр P = 2·(a + b)
package main

import "fmt"

func main() {
	var x, y float64
	fmt.Println("Необходимо найти площадь и периметр прямоугольника")
	x = number("Введите значение стороны 1")
    y = number("Введите значение стороны 2")
    fmt.Println("Площадь равна:\t", x*y)
    fmt.Println("Периметр равен:\t", 2*(x+y))
}

func number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	fmt.Scan(&num)
	return num
}
