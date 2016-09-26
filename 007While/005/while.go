// Дано целое число N (> 0), являющееся некоторой степенью числа 2:
// N = 2^K . Найти целое число K — показатель этой степени

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
		x int
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}
	x = countLoop(n)
	if x == 0 {
		fmt.Println("степень числа 2 : нет")
	} else {
		fmt.Printf("степень числа 2 : %v\n", x)
	}
}

func countLoop(n int) int {
	x := 1
	y := 0
	for x < n {
		x = x * 2
		y++
	}
	if x == n {
		return y
	}
	return 0
}
