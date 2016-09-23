// Дано целое число N (> 0). Последовательность вещественных чисел A K
// определяется следующим образом:
// A 0 = 1,
// A K = ( A K–1 + 1)/ K , K = 1, 2, ... .
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

// A K = ( A K–1 + 1)/ K , K = 1, 2, ...
func foo1(n int) {
	a := 1.0
	for i := 1.0; i <= float64(n); i++ {
		fmt.Printf("A(%v) = %v\n", i, (a+1)/i)
		a = (a + 1) / i
	}
}
