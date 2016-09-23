// Даны целые положительные числа N и K . Найти сумму
// 1^K + 2^K + ... + N^K .
// Чтобы избежать целочисленного переполнения, вычислять слагаемые этой
// суммы с помощью вещественной переменной и выводить результат как
// вещественное число.

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n, k uint
	)

	for n <= 0 {
		n = ioutil.UInteger("N")
	}

	for k <= 0 {
		k = ioutil.UInteger("K")
	}

	fmt.Println("1^K + 2^K + ... + N^K")
	fmt.Println(foo1(n, k))
}

// 1^K + 2^K + ... + N^K
func foo1(n, k uint) float64 {
	var i uint
	x := 1.0
	for i = 2; i <= n; i++ {
		x = x + math.Pow(float64(i), float64(k))
	}
	return x
}
