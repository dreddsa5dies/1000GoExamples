// Дано вещественное число A и целое число N (> 0). Используя один цикл,
// найти сумму
// 1 + A + A^2 + A^3 + ... + A^N
package main

import (
	"fmt"
	"math"

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
	fmt.Println("сумма 1 + A + A^2 + A^3 + ... + A^N")
	fmt.Printf("%v\n", foo1(a, n))
}

func foo1(a float64, n int) float64 {
	var k float64
	k = 1 + a
	for i := 2; i <= n; i++ {
		k = k + math.Pow(a, float64(i))
	}
	return k
}
