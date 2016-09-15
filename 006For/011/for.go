// Дано целое число N (> 0). Найти сумму
// 1 + 1/2 + 1/3 + ... + 1/ N
// (вещественное число)
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

	fmt.Println("1 + 1/2 + 1/3 + ... + 1/ N")
	fmt.Printf("%1.4f\n", foo1(m))
}

func foo1(m int) float64 {
	k := 1.0
	for i := 1.0; i <= float64(m); i++ {
		k = k + 1/i
	}
	return k
}
