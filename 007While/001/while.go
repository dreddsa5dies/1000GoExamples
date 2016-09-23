// Даны положительные числа A и B ( A > B ). На отрезке длины A разме-
// щено максимально возможное количество отрезков длины B (без наложе-
// ний). Не используя операции умножения и деления, найти длину незанятой
// части отрезка A

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		a, b int
	)

	for a <= 0 {
		a = ioutil.Integer("A")
	}

	b = a + 1
	for b > a {
		b = ioutil.Integer("B")
	}
	fmt.Printf("%v\n", countLoop(a, b))
}

func countLoop(a, b int) int {
	x := a
	for x >= 0 {
		x = x - b
	}
	return x + b
}
