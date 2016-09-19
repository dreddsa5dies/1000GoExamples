// Дано вещественное число A и целое число N (> 0). Используя один цикл,
// найти значение выражения
// 1 – A + A^2 – A^3 + ... + (–1)^N · A^N .
// Условный оператор не использовать
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
	fmt.Println("1 – A + A^2 – A^3 + ... + (–1)^N · A^N")
	fmt.Printf("%v\n", foo1(a, n))
}

func foo1(a float64, n int) float64 {
	k := 1.0
	x := 1.0
	for i := 1; i <= n; i++ {
		x = -x
		k = k + x*math.Pow(a, float64(i))
	}
	return k
}
