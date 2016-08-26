// Package util утилиты для решения задач
package util

import "fmt"

// Number ввод числа в stdin
func Number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	_, err := fmt.Scanf("%f", &num)
	if err != nil {
		fmt.Println("Ввод неверных данных")
		panic(err)
	}
	return num
}
