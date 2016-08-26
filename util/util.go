// Package util утилиты для решения задач
package util

import "fmt"

// Number ввод числа в stdin
func Number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	return num
}

// NoNNumber ввод неотрицательного числа в stdin
func NoNNumber(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	for num < 0 {
		fmt.Println("Ввод неверных данных, введите неотрицательное число")
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			panic("Ввод неверных данных")
		}
	}
	return num
}

// NotNullNumber ввод неотрицательного числа в stdin
func NotNullNumber(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	for num == 0 {
		fmt.Println("Ввод неверных данных, введите ненулевое число")
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			panic("Ввод неверных данных")
		}
	}
	return num
}

// ModNumber модуля числа
func ModNumber(num float64) float64 {
	fmt.Printf("Приведение к модулю |%v|\n", num)
	if num < 0 {
		num = -num
	}
	return num
}
