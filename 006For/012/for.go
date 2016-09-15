// Дано целое число N (> 0). Найти произведение
// 1.1 · 1.2 · 1.3 · ...
// ( N сомножителей)
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

	fmt.Println("1.1 · 1.2 · 1.3 · ...")
	fmt.Printf("%1.5f\n", foo1(m))
}

func foo1(m int) float64 {
	k := 1.1
	for i := 1; i < m; i++ {
		k *= (k + 0.1)
	}
	return k
}
