// Дано целое число N (> 1). Последовательность чисел Фибоначчи F K
// определяется следующим образом:
// F 1 = 1,
// F 2 = 1,
// F K = F K–2 + F K–1 , K = 3, 4, ... .
// Проверить, является ли число N числом Фибоначчи. Если является, то вы-
// вести True, если нет — вывести False

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
		b bool
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}

	b = countLoop(n)
	fmt.Printf("является ли число N числом Фибоначчи?\n")
	fmt.Printf("%v\n", b)
}

func countLoop(n int) bool {
	f := 0
	f1 := 1
	f2 := 1
	for f < n {
		f = f2 + f1
		f2 = f1
		f1 = f
	}
	b := (f == n)
	return b
}
