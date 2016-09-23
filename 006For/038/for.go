// Дано целое число N (> 0). Найти сумму 1^N + 2^N–1 + ... + N^1 . Чтобы избе-
// жать целочисленного переполнения, вычислять слагаемые этой суммы с
// помощью вещественной переменной и выводить результат как веществен-
// ное число

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
	)

	for n <= 0 {
		n = ioutil.Integer("N")
	}

	fmt.Println("1^N + 2^N–1 + ... + N^1")
	fmt.Println(foo1(n))
}

// 1^N + 2^N–1 + ... + N^1
func foo1(n int) float64 {
	var x float64
	k := n
	for i := 1.0; i <= float64(n); i++ {
		x = x + math.Pow(i, float64(k))
		k = k - 1
	}
	return x
}
