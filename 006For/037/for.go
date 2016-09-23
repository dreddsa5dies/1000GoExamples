// Дано целое число N (> 0). Найти сумму 1^1 + 2^2 + ... + N^N . Чтобы избе-
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

	fmt.Println("1^1 + 2^2 + ... + N^N")
	fmt.Println(foo1(n))
}

// 1^1 + 2^2 + ... + N^N
func foo1(n int) float64 {
	var x float64
	for i := 1.0; i <= float64(n); i++ {
		x = x + math.Pow(i, i)
	}
	return x
}
