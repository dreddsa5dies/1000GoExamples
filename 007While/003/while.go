// Даны целые положительные числа N и K . Используя только операции
// сложения и вычитания, найти частное от деления нацело N на K , а также
// остаток от этого деления

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n, k, x, y int
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}

	for k <= 0 {
		k = ioutil.Integer("K")
	}
	x, y = countLoop(n, k)
	fmt.Printf("Остаток деления : %v\nЧастное от отделения: %v\n", x, y)
}

func countLoop(n, k int) (x, y int) {
	x = n
	y = -1
	for x >= 0 {
		x = x - k
		y++
	}
	x = x + k
	return x, y
}
