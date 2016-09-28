// Дано целое число N (> 0). С помощью операций деления нацело и взя-
// тия остатка от деления определить, имеется ли в записи числа N цифра «2».
// Если имеется, то вывести True, если нет — вывести False

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
	fmt.Printf("в записи числа N цифра «2»: ")
	k := false
	for _, prX := range x {
		if prX == 2 {
			k = true
		}
	}
	fmt.Printf("%v", k)
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
