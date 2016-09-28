// Дано целое число N (> 0). Используя операции деления нацело и взя-
// тия остатка от деления, найти число, полученное при прочтении числа N
// справа налево

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}
	x := countLoop(n)
	fmt.Printf("число = ")
	for _, prX := range x {
		// хитрость вывода
		fmt.Printf("%v", prX)
	}
	fmt.Println()
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
