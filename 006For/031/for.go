// Дано целое число N (> 0). Последовательность вещественных чисел A K
// определяется следующим образом:
// A 0 = 2,
// A K = 2 + 1/ A K–1 , K = 1, 2, ... .
// Вывести элементы A 1 , A 2 , ..., A N

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
	)

	n = -1
	for n <= 1 {
		n = ioutil.Integer("N")
	}

	fmt.Println("A 1 , A 2 , ..., A N")
	foo1(n)
}

// Последовательность вещественных чисел A K
func foo1(n int) {
	a := 2.0
	for i := 1; i <= n; i++ {
		fmt.Printf("A(%v) = %v\n", i, 2+1/a)
		a = 2 + 1/a
	}
}
