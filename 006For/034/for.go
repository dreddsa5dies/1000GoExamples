// Дано целое число N (> 1). Последовательность вещественных чисел A K
// определяется следующим образом:
// A 1 = 1,
// A 2 = 2,
// A K = ( A K–2 + 2· A K–1 )/3, K = 3, 4, ... .
// Вывести элементы A 1 , A 2 , ..., A N .

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
	fmt.Println("A(1) = 1")
	fmt.Println("A(2) = 2")
	foo1(n)
}

// A K = ( A K–2 + 2· A K–1 )/3, K = 3, 4, ...
func foo1(n int) {
	var a float64
	a1 := 1.0
	a2 := 2.0
	for i := 3; i <= n; i++ {
		a = (a1 + 2*a2) / 3
		fmt.Printf("A(%v) = %1.40f\n", i, a)
		a1 = a2
		a2 = a
	}
}
