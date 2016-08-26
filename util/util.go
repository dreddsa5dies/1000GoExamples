// Package util утилиты для решения задач
package util

import (
	"fmt"
)

// Number ввод числа в stdin
func Number(msg string) float64 {
	fmt.Print(msg + " > ")
	var num float64
	fmt.Scan(&num)
	return num
}