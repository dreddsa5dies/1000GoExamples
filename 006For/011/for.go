// Дано целое число N (> 0). Найти сумму
// N^2 + ( N + 1) 2 + ( N + 2) 2 + ... + (2· N )^2
// (целое число)
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

	fmt.Println("N^2 + ( N + 1)^2 + ( N + 2)^2 + ... + (2· N )^2")
	fmt.Printf("%v\n", foo1(m))
}

func foo1(m int) int {
	k := 0
	for i := 1; i <= m; i++ {
		k += (m + i) * (m + i)
	}
	return k + m*m
}
