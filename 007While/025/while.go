// Дано целое число N (> 1). Найти первое число Фибоначчи, большее N .
// (определение чисел Фибоначчи дано в задании While24)

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

	k = countLoop(n)
	fmt.Printf("является ли число N числом Фибоначчи?\n")
	fmt.Printf("%v\n", k)
}

func countLoop(n int) int {
	f := 0
	f1 := 1
	f2 := 1
	for f < n+1 {
		f = f2 + f1
		f2 = f1
		f1 = f
	}
	return f
}
