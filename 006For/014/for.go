// Дано целое число N (> 0). Найти квадрат данного числа, используя для
// его вычисления следующую формулу:
// N^2 = 1 + 3 + 5 + ... + (2· N – 1)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var m int
	m = -1
	for m < 0 {
		m = ioutil.Integer("N")
	}

	fmt.Println("N^2 = 1 + 3 + 5 + ... + (2· N – 1)")
	fmt.Printf("%v\n", foo1(m))
}

func foo1(m int) int {
	k := 0
	for i := 1; i <= m; i++ {
		k = k + (2*i - 1)
	}
	return k
}
