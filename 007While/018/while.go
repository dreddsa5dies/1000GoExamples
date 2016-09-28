// Дано целое число N (> 0). Используя операции деления нацело и взя-
// тия остатка от деления, найти количество и сумму его цифр

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
		k int
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}
	x := countLoop(n)
	fmt.Printf("количество = %v\n", len(x))
	fmt.Printf("сумму его цифр = ")
	for _, prX := range x {
		k += prX
	}
	fmt.Printf("%v\n", k)
}

func countLoop(n int) []int {
	var y []int
	x := n
	y = append(y, x%10)
	for x > 0 {
		x = x / 10
		y = append(y, x%10)
	}
	k := len(y) - 1
	return y[0:k]
}
