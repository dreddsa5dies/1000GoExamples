// Дано целое число N (> 1), являющееся числом Фибоначчи: N = F K (оп-
// ределение чисел Фибоначчи дано в задании While24). Найти целое число K
// — порядковый номер числа Фибоначчи N

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
	fmt.Printf("порядковый номер числа Фибоначчи N\n")
	fmt.Printf("%v\n", k)
}

func countLoop(n int) int {
	f := 0
	f1 := 1
	f2 := 1
	k := 2
	for f < n {
		k++
		f = f2 + f1
		f2 = f1
		f1 = f
	}
	return k
}
