// Дано целое число N (> 1), являющееся числом Фибоначчи: N = F K (оп-
// ределение чисел Фибоначчи дано в задании While24). Найти целые числа
// F K–1 и F K+1 — предыдущее и последующее числа Фибоначчи

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n      int
		k1, k2 int
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}

	k1, k2 = countLoop(n)
	fmt.Printf("F K–1 и F K+1 — предыдущее и последующее числа Фибоначчи\n")
	fmt.Printf("%v\t%v\n", k1, k2)
}

func countLoop(n int) (x, y int) {
	f := 0
	f1 := 1
	f2 := 1
	for f < n {
		f = f2 + f1
		f2 = f1
		f1 = f
	}
	return f2, f1 + f2
}
