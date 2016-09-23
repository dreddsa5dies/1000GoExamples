// Дано целое число N (> 1). Последовательность чисел Фибоначчи F K
// (целого типа) определяется следующим образом:
// F 1 = 1,
// F 2 = 1,
// F K = F K–2 + F K–1 , K = 3, 4, ... .
// Вывести элементы F 1 , F 2 , ..., F N .

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

	fmt.Println("F 1 , F 2 , ..., F N")
	fmt.Println("F(1) = 1")
	fmt.Println("F(2) = 1")
	foo1(n)
}

// F K = F K–2 + F K–1 , K = 3, 4, ...
func foo1(n int) {
	var fk int
	fk1 := 1
	fk2 := 1
	for i := 3; i <= n; i++ {
		fk = fk2 + fk1
		fmt.Printf("F(%v) = %v\n", i, fk)
		fk2 = fk1
		fk1 = fk
	}
}
