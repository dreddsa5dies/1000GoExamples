// Дано целое число N (> 2). Последовательность целых чисел A K опреде-
// ляется следующим образом:
// A 1 = 1,
// A 2 = 2,
// A 3 = 3,
// A K = A K–1 + A K–2 – 2· A K–3 , K = 4, 5, ... .
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
	for n <= 2 {
		n = ioutil.Integer("N")
	}

	fmt.Println("A 1 , A 2 , ..., A N")
	fmt.Println("A(1) = 1")
	fmt.Println("A(2) = 2")
	fmt.Println("A(3) = 3")
	foo1(n)
}

// A K = A K–1 + A K–2 – 2· A K–3 , K = 4, 5, ...
func foo1(n int) {
	var a int
	a1 := 1
	a2 := 2
	a3 := 3
	for i := 4; i <= n; i++ {
		a = a3 + a2 - 2*a1
		fmt.Printf("A(%v) = %v\n", i, a)
		a1 = a2
		a2 = a3
		a3 = a
	}
}
