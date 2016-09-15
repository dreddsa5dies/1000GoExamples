// Дано вещественное число A и целое число N (> 0). Найти A в степени N :
// A^N = A · A · ... · A
// (числа A перемножаются N раз)
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		a float64
		n int
	)

	a = ioutil.Number("A")

	n = -1
	for n < 0 {
		n = ioutil.Integer("N")
	}
	fmt.Println("A^N = A · A · ... · A")
	fmt.Printf("%v\n", foo1(a, n))
}

func foo1(a float64, n int) float64 {
	k := 1.0
	for i := 1; i <= n; i++ {
		k = k * a
	}
	return k
}
