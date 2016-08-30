// Package util утилиты для решения задач
package util

import (
	"fmt"
	"math"
)

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

// NotNullNumber ввод ненулевого числа в stdin
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

// GrNumber ввод градусов (0 < α < 360) в stdin
func GrNumber(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	for num < 0 || num > 360 {
		fmt.Println("Ввод неверных данных, введите 0 < α < 360")
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			panic("Ввод неверных данных")
		}
	}
	return num
}

// RadNumber ввод радиан (0 < α < 2·π) в stdin
func RadNumber(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		panic("Ввод неверных данных")
	}
	for num < 0 || num > 2*math.Pi {
		fmt.Println("Ввод неверных данных, введите 0 < α < 2·π")
		_, err := fmt.Scanf("%f", &num)
		if err != nil {
			panic("Ввод неверных данных")
		}
	}
	return num
}

// ModNumber приведение к модулю числа
func ModNumber(num float64) float64 {
	if num < 0 {
		num = -num
	}
	return num
}
